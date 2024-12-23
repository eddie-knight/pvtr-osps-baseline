package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All proposed changes to the project’s codebase must be automatically evaluated against a documented policy for known vulnerabilities and blocked in the event of violations except when declared and supressed as non exploitable.",
		ControlID:   "OSPS-QA-07",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(QA_07_T01)

	return "QA_07", result
}

// TODO
func QA_07_T01() pluginkit.TestResult {
	moveResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_07
	return moveResult
}
