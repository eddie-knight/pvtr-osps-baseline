package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_05() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project’s permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated.",
		ControlID:   "OSPS-AC-05",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_05_T01)

	return "AC_05", result
}

func AC_05_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return testResult
}
