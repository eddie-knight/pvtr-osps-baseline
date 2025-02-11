package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func VM_05() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project publishes contacts and process for reporting vulnerabilities.",
		ControlID:   "OSPS-VM-05",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(VM_05_T01)

	return "VM_05", result
}

// TODO
func VM_05_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to VM_05
	return testResult
}
