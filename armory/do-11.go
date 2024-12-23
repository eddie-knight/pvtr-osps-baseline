package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_11() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources.",
		ControlID:   "OSPS-DO-11",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_11_T01)

	return "DO_11", result
}

// TODO
func DO_11_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_11
	return moveResult
}
