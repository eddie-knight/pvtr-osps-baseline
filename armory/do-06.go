package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_06() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions.",
		ControlID:   "OSPS-DO-06",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_06_T01)
	return "DO_06", result
}

func DO_06_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_06
	return testResult
}
