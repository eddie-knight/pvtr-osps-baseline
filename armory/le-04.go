package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"

	"github.com/privateerproj/privateer-sdk/raidengine"
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
	result.ExecuteTest(LE_04_T02)
	result.ExecuteTest(LE_04_T03)

	return "LE_04", result
}

func LE_04_T01() raidengine.MovementResult {
	releases := Data.GraphQL().Repository.Releases
	if releases.TotalCount == 0 {
		return raidengine.MovementResult{
			Description: "Check release license compliance",
			Function:    utils.CallerPath(0),
			Passed:      false,
			Message:     "No releases found in repository",
		}
	}

	return raidengine.MovementResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Message:     fmt.Sprintf("Found %d releases", releases.TotalCount),
	}
}

func LE_04_T02() raidengine.MovementResult {
	releases := Data.GraphQL().Repository.Releases
	if releases.TotalCount == 0 {
		return raidengine.MovementResult{
			Description: "Check release license compliance",
			Function:    utils.CallerPath(0),
			Passed:      false,
			Message:     "No releases to check for license",
		}
	}

	latestRelease := releases.Nodes[0]
	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			return raidengine.MovementResult{
				Description: "Check release license compliance",
				Function:    utils.CallerPath(0),
				Passed:      true,
				Message:     fmt.Sprintf("Found license file: %s", asset.Name),
			}
		}
	}

	return raidengine.MovementResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     fmt.Sprintf("No license file found in release %s", latestRelease.TagName),
	}
}

func LE_04_T03() raidengine.MovementResult {
	releases := Data.GraphQL().Repository.Releases
	if releases.TotalCount == 0 {
		return raidengine.MovementResult{
			Description: "Check release license compliance",
			Function:    utils.CallerPath(0),
			Passed:      false,
			Message:     "No releases to validate license",
		}
	}

	latestRelease := releases.Nodes[0]
	for _, asset := range latestRelease.ReleaseAssets.Nodes {
		if strings.Contains(strings.ToLower(asset.Name), "license") {
			if asset.Content == "" {
				return raidengine.MovementResult{
					Description: "Check release license compliance",
					Function:    utils.CallerPath(0),
					Passed:      false,
					Message:     fmt.Sprintf("Empty license content in release %s", latestRelease.TagName),
				}
			}

			upperContent := strings.ToUpper(asset.Content)
			for spdxId := range approvedSpdx {
				if strings.Contains(upperContent, spdxId) {
					return raidengine.MovementResult{
						Description: "Check release license compliance",
						Function:    utils.CallerPath(0),
						Passed:      true,
						Message:     fmt.Sprintf("Valid SPDX license found: %s", spdxId),
					}
				}
			}

			return raidengine.MovementResult{
				Description: "Check release license compliance",
				Function:    utils.CallerPath(0),
				Passed:      false,
				Message:     "License content does not match approved SPDX identifiers",
			}
		}
	}

	return raidengine.MovementResult{
		Description: "Check release license compliance",
		Function:    utils.CallerPath(0),
		Passed:      false,
		Message:     "No license file found to validate",
	}
}
