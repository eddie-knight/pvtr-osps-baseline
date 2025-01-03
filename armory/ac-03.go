package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project's version control system MUST prevent unintentional direct commits against the primary branch.",
		ControlID:   "OSPS-AC-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(AC_03_T01)

	return "AC_03", result
}

func AC_03_T01() pluginkit.TestResult {
	protectionData := Data.GraphQL().Repository.DefaultBranchRef.BranchProtectionRule
	// TODO: check rulesets also

	var message string
	if protectionData.RestrictsPushes {
		message = "Branch protection rule restricts pushes"
	} else if protectionData.RequiresApprovingReviews {
		message = "Branch protection rule requires approving reviews"
	}

	testResult := pluginkit.TestResult{
		Description: "Inspect default branch for a protection rule that restricts pushes",
		Function:    utils.CallerPath(0),
		Passed:      protectionData.RestrictsPushes || protectionData.RequiresApprovingReviews,
		Message:     message,
	}

	return testResult
}
