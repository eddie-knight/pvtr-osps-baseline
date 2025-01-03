package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_02() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project’s version control system MUST restrict collaborator permissions to the lowest available privileges by default.",
		ControlID:   "OSPS-AC-02",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_02_T01)

	return "AC_02", result
}

func AC_02_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return testResult
}
