package armory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ossf/si-tooling/v2/si"
)

type RestData struct {
	owner    string
	repo     string
	Repo     RepoData
	Insights si.SecurityInsights
}

type RepoData struct {
	Name     string `json:"name"`
	Private  bool   `json:"private"`
	Releases []ReleaseData
	Contents struct {
		TopLevel []DirContents
		ForgeDir []DirContents
	}
}

type ReleaseData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	TagName string `json:"tag_name"`
}

type DirContents struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	SHA         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

type FileAPIResponse struct {
	ByteContent []byte `json:"content"`
	SHA         string `json:"sha"`
}

var APIBase = "https://api.github.com/repos"

func makeApiCall(endpoint string, authRequired bool) (body []byte, err error) {
	Logger.Trace(fmt.Sprintf("GET %s", endpoint))
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if authRequired && Authenticated {
		request.Header.Set("Authorization", "Bearer "+GlobalConfig.GetString("token"))
	} else if authRequired && !Authenticated {
		err = fmt.Errorf("auth required but not authenticated")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("error making http call: %s", err.Error())
		return nil, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("unexpected response: %s", response.Status)
		return nil, err
	}
	return io.ReadAll(response.Body)
}

func getSourceFile(owner, repo, path string) (response FileAPIResponse, err error) {
	endpoint := fmt.Sprintf("%s/%s/%s/contents/%s", APIBase, owner, repo, path)
	responseData, err := makeApiCall(endpoint, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseData, &response)
	return
}

func (r *RestData) loadData() error {
	r.owner = GlobalConfig.GetString("owner")
	r.repo = GlobalConfig.GetString("repo")

	r.getMetadata()
	if r.Repo.Releases == nil {
		r.Repo.getReleases(r.owner, r.repo)
	}
	r.loadSecurityInsights()
	return nil
}

func (r *RestData) loadSecurityInsights() {
	r.getTopDirContents()
	if len(r.Repo.Contents.TopLevel) == 0 {
		Logger.Error("no contents retrieved from the top level of the repository")
		return
	}
	for _, content := range r.Repo.Contents.TopLevel {
		if r.foundSecurityInsights(content) {
			insights, err := si.Read(r.owner, r.repo, "security-insights.yml")
			r.Insights = insights
			if err != nil {
				Logger.Error(fmt.Sprintf("error reading security insights file: %s", err.Error()))
			}
			return
		}
	}
	r.getForgeDirContents()
	for _, content := range r.Repo.Contents.ForgeDir {
		if r.foundSecurityInsights(content) {
			insights, err := si.Read(r.owner, r.repo, ".github/security-insights.yml")
			r.Insights = insights
			if err != nil {
				Logger.Error(fmt.Sprintf("error reading security insights file: %s", err.Error()))
			}
			return
		}
	}
	Logger.Error("no security insights file found")
}

func (r *RestData) foundSecurityInsights(content DirContents) bool {
	if strings.Contains(strings.ToLower(content.Name), "security-insights.") {
		response, err := getSourceFile(r.owner, r.repo, content.Path)
		if err != nil {
			Logger.Error(fmt.Sprintf("error unmarshalling API response for security insights file: %s", err.Error()))
			return false
		}
		Logger.Trace(fmt.Sprintf("Security Insights Exists - SHA: %v", response.SHA))
		return true
	}
	return false
}

func (r *RestData) getTopDirContents() {
	endpoint := fmt.Sprintf("%s/%s/%s/contents", APIBase, r.owner, r.repo)
	responseData, err := makeApiCall(endpoint, false)
	if err != nil {
		Logger.Error(fmt.Sprintf("error getting top level contents: %s", err.Error()))
		return
	}
	json.Unmarshal(responseData, &r.Repo.Contents.TopLevel)
}

func (r *RestData) getForgeDirContents() {
	endpoint := fmt.Sprintf("%s/%s/%s/contents/.github", APIBase, r.owner, r.repo)
	responseData, err := makeApiCall(endpoint, false)
	if err != nil {
		Logger.Error(fmt.Sprintf("error getting forge contents: %s", err.Error()))
		return
	}
	json.Unmarshal(responseData, &r.Repo.Contents.ForgeDir)
}

func (r *RestData) getMetadata() error {
	endpoint := fmt.Sprintf("%s/%s/%s", APIBase, r.owner, r.repo)
	responseData, err := makeApiCall(endpoint, false)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, &r.Repo)
}

func (r *RepoData) getReleases(owner, repo string) error {
	endpoint := fmt.Sprintf("%s/%s/%s/releases", APIBase, owner, repo)
	responseData, err := makeApiCall(endpoint, false)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, &r.Releases)
}

func (r *RestData) GetFileContentByURL(downloadURL string) (string, error) {
	// Call the same low-level function used by the rest of your data-rest flows
	responseData, err := makeApiCall(downloadURL, true)
	if err != nil {
		return "", err
	}
	// Convert the raw bytes to a string and return
	return string(responseData), nil
}
