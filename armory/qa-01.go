package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's source code MUST be publicly readable and have a static URL.",
		ControlID:   "OSPS-QA-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(QA_01_T01)

	return "QA_01", result
}

func QA_01_T01() pluginkit.TestResult {
	gotRepoData := Data.Rest().Repo.Name != ""
	isPrivate := Data.Rest().Repo.Private

	testResult := pluginkit.TestResult{
		Description: "Verifying that the GitHub repository is public at the target URL.",
		Function:    utils.CallerPath(0),
		Passed:      gotRepoData && !isPrivate,
	}

	if !gotRepoData {
		testResult.Message = "Repository data not found"
	} else {
		testResult.Message = fmt.Sprintf("Public Repo: %t", !isPrivate)
	}

	return testResult
}
