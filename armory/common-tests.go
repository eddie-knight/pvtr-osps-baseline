package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func countReleases() pluginkit.TestResult {
	releases := Data.Rest().Repo.Releases

	testResult := pluginkit.TestResult{
		Description: "Counting the number of releases on the repository",
		Function:    utils.CallerPath(1),
		Passed:      true,
		Value:       len(releases),
		Message:     fmt.Sprintf("Releases found: %v", len(releases)),
	}

	return testResult
}
