package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_05() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated.",
		ControlID:   "OSPS-AC-05",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_05_T01)

	return "AC_05", result
}

func AC_05_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "GitHub Actions workflow permissions must be configured with minimal access",
		Function:    utils.CallerPath(0),
		Passed:      true, // default pass unless violations found
	}

	rest := Data.Rest()
	permResp, err := rest.getWorkflowPermissions()

	if err != nil {
		testResult.Message = err.Error()
		testResult.Passed = false
		return testResult
	}

	// Check default workflow permissions
	if permResp.DefaultWorkflowPermissions != "read" {
		testResult.Passed = false
		testResult.Message = fmt.Sprintf("default workflow permissions are not minimal (current: %s)", permResp.DefaultWorkflowPermissions)
		return testResult
	}

	return testResult
}
