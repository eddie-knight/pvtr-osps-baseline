package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST provide user guides for all basic functionality.",
		ControlID:   "OSPS-DO-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_03_T01)

	return "DO_03", result
}

func DO_03_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: This needs security insights data
	return moveResult
}
