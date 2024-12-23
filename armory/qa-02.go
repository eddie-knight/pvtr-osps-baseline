package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_02() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made.",
		ControlID:   "OSPS-QA-02",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(QA_02_T01)

	return "QA_02", result
}

// TODO
func QA_02_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_02
	return moveResult
}
