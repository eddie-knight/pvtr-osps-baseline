package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The license for the source code MUST be maintained in a standard location within the project’s repository.",
		ControlID:   "OSPS-LE-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_03_T01)

	return "LE_03", result
}

// TODO
func LE_03_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_03
	return testResult
}
