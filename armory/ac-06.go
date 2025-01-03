package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_06() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project’s version control system MUST require multi-factor authentication that does not include SMS for users when modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-06",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_06_T01)

	return "AC_06", result
}

func AC_06_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return testResult
}
