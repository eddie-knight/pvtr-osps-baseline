package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All released software assets MUST be delivered with a machine-readable list of all direct and transitive internal software dependencies with their associated version identifiers.",
		ControlID:   "OSPS-QA-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(QA_03_T01)

	return "QA_03", result
}

// TODO
func QA_03_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_03
	return moveResult
}
