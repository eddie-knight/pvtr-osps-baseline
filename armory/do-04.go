package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_04() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response.",
		ControlID:   "OSPS-DO-04",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_04_T01)

	return "DO_04", result
}

func DO_04_T01() pluginkit.TestResult {
	enabled := Data.GraphQL().Repository.IsSecurityPolicyEnabled

	testResult := pluginkit.TestResult{
		Description: "Discover whether a security policy is enabled for this repo.",
		Function:    utils.CallerPath(0),
		Passed:      enabled,
		Message:     fmt.Sprintf("Security Policy Enabled: %v", enabled),
	}
	return testResult
}
