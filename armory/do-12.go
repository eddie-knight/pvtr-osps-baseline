package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_12() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST contain instructions to verify the integrity and authenticity of the release assets, including the expected identity of the person or process authoring the software release.",
		ControlID:   "OSPS-DO-12",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_12_T01)

	return "DO_12", result
}

// TODO
func DO_12_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_12
	return moveResult
}
