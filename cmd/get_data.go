package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/privateerproj/privateer-raid-osps-baseline/armory"
	"github.com/privateerproj/privateer-sdk/config"
)

func getDataFromRestApi(c *config.Config) ([]byte, error) {
	repo := c.GetString("repo")
	owner := c.GetString("owner")
	token := c.GetString("auth-token")

	if token != "" {
		// do something different if authenticated
		return nil, nil
	} else {
		response, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo))
		if err != nil {
			err = fmt.Errorf("error making http call: %s", err.Error())
			return nil, err
		}
		return io.ReadAll(response.Body)
	}
}

func getRepositoryData(c *config.Config) (repoData armory.RepoData, err error) {
	responseData, err := getDataFromRestApi(c)
	if err != nil {
		return
	}
	json.Unmarshal(responseData, &repoData)
	return
}
