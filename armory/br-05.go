package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_05() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All build and release pipelines MUST use standardized tooling where available to ingest dependencies at build time.",
		ControlID:   "OSPS-BR-05",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_05_T01)

	return "BR_05", result
}

func BR_05_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return testResult
}
