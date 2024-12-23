package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project’s build and release pipelines MUST NOT execute arbitrary code that is input from outside of the build script.",
		ControlID:   "OSPS-BR-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_01_T01)

	return "BR_01", result
}

// TODO
func BR_01_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}
