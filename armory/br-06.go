package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_06() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All releases MUST provide a descriptive log of functional and security modifications.",
		ControlID:   "OSPS-BR-06",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_06_T01)
	if !strings.Contains(result.Tests["BR_06_T01"].Message, ": 0") {
		Logger.Trace("Releases Found, checking for Change Log")
		result.ExecuteTest(BR_06_T02)
	}
	return "BR_06", result
}

func BR_06_T01() pluginkit.TestResult {
	releaseCount := Data.GraphQL().Repository.Releases.TotalCount

	return pluginkit.TestResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Message:     fmt.Sprintf("Releases Found: %v", releaseCount),
	}
}

func BR_06_T02() pluginkit.TestResult {
	releaseDescription := Data.GraphQL().Repository.LatestRelease.Description
	contains := (strings.Contains(releaseDescription, "Change Log") || strings.Contains(releaseDescription, "Changelog"))

	testResult := pluginkit.TestResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
		Passed:      contains,
		Message:     fmt.Sprintf("Change Log Found in Latest Release: %v", contains),
	}
	return testResult
}
