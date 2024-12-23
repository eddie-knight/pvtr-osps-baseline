package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "Any websites, API responses or other services involved in the project development and release MUST be delivered using SSH, HTTPS or other encrypted channels.",
		ControlID:   "OSPS-BR-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_03_T01)

	return "BR_03", result
}

func BR_03_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}
