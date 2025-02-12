package armory

import (
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_09() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "Any websites or other services involved in the distribution of released software assets MUST be delivered using HTTPS or other encrypted channels.",
		ControlID:   "OSPS-BR-09",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_09_T01)

	return "BR_09", result
}

func BR_09_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "Ensure that no URIs in the Security Insights Release Distribution points are using HTTP.",
		Function:    utils.CallerPath(0),
		Passed:      true,
	}

	links := Data.Rest().Insights.Repository.Release.DistributionPoints
	badLinks := []string{}
	for _, link := range links {
		if strings.Contains(link.URI, "http://") {
			testResult.Passed = false
			badLinks = append(badLinks, link.URI)
		}
	}

	if !testResult.Passed {
		testResult.Value = badLinks
		testResult.Message = "Found one or more distribution point using HTTP instead of HTTPS."
	} else {
		testResult.Message = "All distribution points are using HTTPS."
	}

	return testResult
}
