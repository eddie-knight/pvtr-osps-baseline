package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_06() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The version control system MUST NOT contain generated executable artifacts.",
		ControlID:   "OSPS-QA-06",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(QA_06_T01)

	return "QA_06", result
}

func QA_06_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_06
	return moveResult
}
