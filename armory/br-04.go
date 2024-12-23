package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_04() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All released software assets MUST be created with consistent, automated build and release pipelines.",
		ControlID:   "OSPS-BR-04",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_04_T01)

	return "BR_04", result
}

func BR_04_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}
