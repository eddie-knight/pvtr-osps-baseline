package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's version control system MUST require multi-factor authentication that does not include SMS for users when modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-07",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_02_T01)

	return "AC_07", result
}

func AC_07_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "Ensure that the project settings require multi-factor authentication that does not include SMS for users with access to sensitive elements.",
		Message:     "GitHub does not allow enforcement of non-SMS MFA.",
		Passed:      false,
		Function:    utils.CallerPath(0),
	}

	return testResult
}
