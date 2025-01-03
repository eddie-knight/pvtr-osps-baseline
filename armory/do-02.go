package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_02() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include an explanation of the contribution process.",
		ControlID:   "OSPS-DO-02",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_02_T01)

	return "DO_02", result
}

func DO_02_T01() pluginkit.TestResult {

	body := Data.GraphQL().Repository.ContributingGuidelines.Body
	containsGuidelines := len(body) > 100

	testResult := pluginkit.TestResult{
		Description: "Discover whether the GitHub repo's recommended contributing guidelines has content.",
		Function:    utils.CallerPath(0),
		Passed:      containsGuidelines,
		Message:     fmt.Sprintf("Contributing Guidelines Found: %v", containsGuidelines),
	}
	return testResult
}
