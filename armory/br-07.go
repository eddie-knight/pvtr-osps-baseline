package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All released software assets MUST be signed or accounted for in a signed manifest including each asset’s cryptographic hashes.",
		ControlID:   "OSPS-BR-07",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_07_T01)

	return "BR_07", result
}

func BR_07_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return testResult
}
