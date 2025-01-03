package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_01_T01)

	return "AC_01", result
}

// TODO
func AC_01_T01() pluginkit.TestResult {
	required := Data.GraphQL().Organization.RequiresTwoFactorAuthentication

	testResult := pluginkit.TestResult{
		Description: "Inspect the repo's parent to ensure that all members are required to use MFA",
		Function:    utils.CallerPath(0),
		Passed:      required,
		Message:     fmt.Sprintf("MFA Required: %v", required),
	}

	return testResult
}
