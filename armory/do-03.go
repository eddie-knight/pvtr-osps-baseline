package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST provide user guides for all basic functionality.",
		ControlID:   "OSPS-DO-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_03_T01)
	if !result.Tests["DO_03_T01"].Passed {
		// TODO: call the location and check for contents
	}
	return "DO_03", result
}

func DO_03_T01() pluginkit.TestResult {
	guideLocation := Data.Rest().Insights.Project.Documentation.DetailedGuide
	found := guideLocation != ""

	moveResult := pluginkit.TestResult{
		Description: "Determine whether the project's Security Insights data specifies a documentation location.",
		Function:    utils.CallerPath(0),
		Passed:      found,
		Message:     fmt.Sprintf("Detailed Guide documentation location specified in Security Insights: %t", found),
		Value:       guideLocation,
	}

	return moveResult
}
