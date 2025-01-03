package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST provide design documentation demonstrating all actions and actors within the system.",
		ControlID:   "OSPS-AC-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_07_T01)

	return "DO_07", result
}

// TODO
func DO_07_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_07
	return testResult
}
