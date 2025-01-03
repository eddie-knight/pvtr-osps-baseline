package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_08() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a policy that defines a threshold for remediation of SCA findings related to vulnerabilities and licenses.",
		ControlID:   "OSPS-DO-08",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_08_T01)

	return "DO_08", result
}

// TODO
func DO_08_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_08
	return testResult
}
