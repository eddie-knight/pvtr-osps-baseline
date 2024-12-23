package armory

import (
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles.",
		ControlID:   "OSPS-DO-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(DO_01_T01)

	return "DO_01", result
}

// TODO
func DO_01_T01() pluginkit.TestResult {
	repoData := Data.GraphQL().Repository

	var message string
	if repoData.HasDiscussionsEnabled {
		message = "Discussions are enabled. "
	}
	if repoData.HasIssuesEnabled {
		message = message + "Issues are enabled."
	}

	moveResult := pluginkit.TestResult{
		Description: "Discover whether issues or discussions are enabled on the repo.",
		Function:    utils.CallerPath(0),
		Message:     message,
		Passed:      repoData.HasDiscussionsEnabled || repoData.HasIssuesEnabled,
	}
	return moveResult
}
