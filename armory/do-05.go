package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_05() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a mechanism for reporting defects.",
		ControlID:   "OSPS-DO-05",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_05_T01)

	return "DO_05", result
}

func DO_05_T01() pluginkit.TestResult {
	enabled := Data.GraphQL().Repository.HasIssuesEnabled
	moveResult := pluginkit.TestResult{
		Description: "Checking to see whether the GitHub repo has issues enabled",
		Function:    utils.CallerPath(0),
		Passed:      enabled,
		Message:     fmt.Sprintf("Issues Enabled: %v", enabled),
	}

	return moveResult
}
