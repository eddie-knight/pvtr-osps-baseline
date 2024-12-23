package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_10() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a policy to address SCA violations prior to any release.",
		ControlID:   "OSPS-DO-10",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_10_T01)

	return "DO_10", result
}

// TODO
func DO_10_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_10
	return moveResult
}
