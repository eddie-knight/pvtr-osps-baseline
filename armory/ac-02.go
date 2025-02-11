package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_02() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's version control system MUST restrict collaborator permissions to the lowest available privileges by default.",
		ControlID:   "OSPS-AC-02",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_02_T01)

	return "AC_02", result
}

func AC_02_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "Ensure the project uses a version control system that restricts collaborator permissions to the lowest available privileges by default.",
		Function:    utils.CallerPath(0),
	}

	dataFound := Data.Rest().Repo.Name != ""
	if !dataFound {
		testResult.Message = "GitHub restricts collaborator permissions to the lowest available privileges by default."
		testResult.Passed = true
	} else {
		testResult.Message = "GitHub repo data not retrieved."
		testResult.Passed = false
	}

	return testResult
}
