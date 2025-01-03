package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_06() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions.",
		ControlID:   "OSPS-DO-06",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_06_T01)
	return "DO_06", result
}

func DO_06_T01() pluginkit.TestResult {
	guideLocation := Data.Rest().Insights.Repository.Documentation.Contributing
	found := guideLocation != ""
	testResult := pluginkit.TestResult{
		Description: "Ensure the project's Security Insights data specifies a contributing guide location.",
		Function:    utils.CallerPath(0),
		Passed:      found,
		Message:     fmt.Sprintf("Contributing guide location specified in Security Insights: %v", found),
		Value:       guideLocation,
	}

	return testResult
}
