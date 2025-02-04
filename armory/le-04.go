package armory

import (
	"fmt"
	"slices"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

// TODO: this is common between LE-03 and LE-04
var approvedLicenses = []string{
	"MIT",
	"GPL-2.0",
	"GPL-3.0",
	"BSD-2-CLAUSE",
	"BSD-3-CLAUSE",
	"APACHE-2.0",
	"LGPL-2.1",
	"LGPL-3.0",
	"MPL-2.0",
	"EPL-2.0",
}

func LE_04() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-04",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_04_T01)
	if v, ok := result.Tests["LE_04_T01"].Value.(int); ok && v > 0 {
		result.ExecuteTest(LE_04_T02)
	}
	if v, ok := result.Tests["LE_04_T02"].Value.(bool); ok && v {
		result.ExecuteTest(LE_04_T03)
	}

	return "LE_04", result
}

func LE_04_T01() pluginkit.TestResult {
	return countReleases()
}

func LE_04_T02() pluginkit.TestResult {
	// Set up the default result
	testResult := pluginkit.TestResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Message:     "No license file found",
		Value:       false,
	}

	latestRelease := Data.Rest().Repo.Releases[0]

	testResult.Message = fmt.Sprintf("No license file found in release %s", latestRelease.TagName)

	for _, asset := range latestRelease.Assets {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			testResult.Message = fmt.Sprintf("Found license file: %s", asset.Name)
			testResult.Value = true
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
		Message:     "No license detected in the latest release assets",
		Passed:      true, // Assume this passes until a bad license is found
	}

	latestRelease := Data.Rest().Repo.Releases[0]

	foundValue := []string{}
	for _, asset := range latestRelease.Assets {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			foundValue = append(foundValue, asset.Name)
			content, err := getFileContentByURL(asset.DownloadURL)
			if err != nil {
				testResult.Message = fmt.Sprintf("Found an apparent license entry, but failed to fetch content: %v", err)
				break
			}
			upperContent := strings.ToUpper(content)
			if !slices.Contains(approvedLicenses, upperContent) {
				testResult.Passed = false
				testResult.Message = fmt.Sprintf("Invalid license found in release assets: %s", asset.Name)
			}
		}
	}
	testResult.Value = foundValue
	return testResult
}
