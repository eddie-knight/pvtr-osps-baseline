package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

var approvedSpdx = map[string]bool{
	"MIT":          true,
	"GPL-2.0":      true,
	"GPL-3.0":      true,
	"BSD-2-CLAUSE": true,
	"BSD-3-CLAUSE": true,
	"APACHE-2.0":   true,
	"LGPL-2.1":     true,
	"LGPL-3.0":     true,
	"MPL-2.0":      true,
	"EPL-2.0":      true,
}

func LE_04() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-04",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_04_T01)
	if count, ok := result.Tests["LE_04_T01"].Value.(int); ok && count > 0 {
		result.ExecuteTest(LE_04_T02)
	}
	if result.Tests["LE_04_T02"].Passed {
		result.ExecuteTest(LE_04_T03)
	}

	return "LE_04", result
}

func LE_04_T01() pluginkit.TestResult {
	releases := Data.GraphQL().Repository.Releases
	return pluginkit.TestResult{
		Description: "Check for available releases",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Message:     fmt.Sprintf("Found %d releases", releases.TotalCount),
		Value:       releases.TotalCount,
	}
}

func LE_04_T02() pluginkit.TestResult {
	// Set up the default result
	testResult := pluginkit.TestResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     "No license file found",
	}

	releases := Data.GraphQL().Repository.Releases
	latestRelease := releases.Nodes[0]

	// Adjust our default message to reflect the actual release
	testResult.Message = fmt.Sprintf("No license file found in release %s", latestRelease.TagName)

	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			testResult.Passed = true
			testResult.Message = fmt.Sprintf("Found license file: %s", asset.Name)
			break
		}
	}

	return testResult
}

func LE_04_T03() pluginkit.TestResult {
	// Default test result
	testResult := pluginkit.TestResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     "No valid license found in release assets",
	}

	releases := Data.GraphQL().Repository.Releases
	latestRelease := releases.Nodes[0]

	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			rest := Data.Rest()
			content, err := rest.GetFileContentByURL(asset.DownloadUrl)
			if err != nil {
				testResult.Message = fmt.Sprintf("Failed to fetch license content: %v", err)
				break // We can break or return early. Here we break so we can still return testResult below.
			}

			upperContent := strings.ToUpper(content)
			for spdxId := range approvedSpdx {
				if strings.Contains(upperContent, spdxId) {
					testResult.Passed = true
					testResult.Message = fmt.Sprintf("Valid SPDX license found: %s", spdxId)
					break
				}
			}

			// If we found it, no need to keep checking other assets
			if testResult.Passed {
				break
			}
		}
	}

	return testResult
}
