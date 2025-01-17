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
		Value:       releases,
		Message:     fmt.Sprintf("Releases found: %v", len(releases)),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return testResult
}
