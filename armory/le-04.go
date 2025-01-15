package armory

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

var approvedSpdx = map[string]bool{
	"MIT":          true,
	"GPL-2.0":      true,
	"GPL-3.0":      true,
	"BSD-2-CLAUSE": true,
	"BSD-3-CLAUSE": true,
	"APACHE-2.0":   true,
	"LGPL-2.1":     true,
	"LGPL-3.0":     true,
	"MPL-2.0":      true,
	"EPL-2.0":      true,
}

func LE_04() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-04",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_04_T01)
	if result.Tests["LE_04_T01"].Value > 0 {
            result.ExecuteTest(LE_04_T02)
	}
	if result.Tests["LE_04_T02"].Passed {
	    result.ExecuteTest(LE_04_T03)
        }

	return "LE_04", result
}

func LE_04_T01() pluginkit.TestResult {
	releases := Data.GraphQL().Repository.Releases
	return pluginkit.TestResult{
		Description: "Check for available releases",
		Function: utils.CallerPath(0),
		Passed: true,
		Message: fmt.Sprintf("Found %d releases", releases.TotalCount),
		Value: releases.TotalCount
	}
}

func LE_04_T02() pluginkit.TestResult {
	releases := Data.GraphQL().Repository.Releases
	if releases.TotalCount == 0 {
		return pluginkit.TestResult{
			Description: "Check release license compliance",
			Function:    utils.CallerPath(0),
			Passed:      false,
			Message:     "No releases to check for license",
		}
	}

	latestRelease := releases.Nodes[0]
	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			return pluginkit.TestResult{
				Description: "Check release license compliance",
				Function:    utils.CallerPath(0),
				Passed:      true,
				Message:     fmt.Sprintf("Found license file: %s", asset.Name),
			}
		}
	}

	return pluginkit.TestResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     fmt.Sprintf("No license file found in release %s", latestRelease.TagName),
	}
}

func LE_04_T03() pluginkit.TestResult {
	releases := Data.GraphQL().Repository.Releases
	if releases.TotalCount == 0 {
		return pluginkit.TestResult{
			Description: "Check release license compliance",
			Function:    utils.CallerPath(0),
			Passed:      false,
			Message:     "No releases to validate license",
		}
	}

	latestRelease := releases.Nodes[0]
	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			content, err := fetchLicenseContent(asset.DownloadUrl)
			if err != nil {
				return pluginkit.TestResult{
					Description: "Check release license compliance",
					Function:    utils.CallerPath(0),
					Passed:      false,
					Message:     fmt.Sprintf("Failed to fetch license content: %v", err),
				}
			}

			upperContent := strings.ToUpper(content)
			for spdxId := range approvedSpdx {
				if strings.Contains(upperContent, spdxId) {
					return pluginkit.TestResult{
						Description: "Check release license compliance",
						Function:    utils.CallerPath(0),
						Passed:      true,
						Message:     fmt.Sprintf("Valid SPDX license found: %s", spdxId),
					}
				}
			}
		}
	}

	return pluginkit.TestResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     "No valid license found in release assets",
	}
}

func fetchLicenseContent(url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	token := GlobalConfig.GetString("token")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("error making http call: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("unexpected response: %s", response.Status)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
