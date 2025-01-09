package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The license for the source code MUST be maintained in a standard location within the project’s repository.",
		ControlID:   "OSPS-LE-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_03_T01)

	return "LE_03", result
}

func LE_03_T01() pluginkit.TestResult {
	license := Data.graphql.Repository.LicenseInfo
	licensePath := Data.graphql.Repository.LicenseInfo.Url

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

	// Check if license exists and is approved
	hasValidLicense := license.Name != "" &&
		license.SpdxId != "" &&
		approvedSpdx[strings.ToUpper(license.SpdxId)]

	standardPaths := []string{
		"LICENSE",
		"LICENSE.md",
		"LICENSE.txt",
		"COPYING",
		"COPYING.md",
		"COPYING.txt",
		"LICENSE/",
	}

	hasStandardPath := false
	for _, path := range standardPaths {
		if strings.Contains(licensePath, path) {
			hasStandardPath = true
			break
		}
	}

	moveResult := pluginkit.TestResult{
		Description: "Verify license is present in a standard location",
		Function:    utils.CallerPath(0),
		Passed:      hasValidLicense && hasStandardPath,
	}

	if !hasValidLicense {
		moveResult.Message = fmt.Sprintf("No valid SPDX license found. Current license: %s (SPDX: %s)",
			license.Name, license.SpdxId)
		return moveResult
	}

	if !hasStandardPath {
		moveResult.Message = fmt.Sprintf("License found but not in standard location. License: %s (SPDX: %s), Path: %s",
			license.Name, license.SpdxId, licensePath)
		return moveResult
	}

	moveResult.Message = fmt.Sprintf("Valid SPDX license found in standard location: %s (SPDX: %s)",
		license.Name, license.SpdxId)

	return moveResult
}
