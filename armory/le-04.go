package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_04() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_04_T01)

	return "LE_04", result
}

// TODO
func LE_04_T01() raidengine.MovementResult {
	license := Data.GraphQL().Repository.LicenseInfo.Name
	Logger.Debug("License::", license)
	// List of common OSI/FSF approved licenses
	approvedLicenses := []string{
		"MIT License",
		"GNU General Public License v2.0",
		"GNU General Public License v3.0",
		"BSD 2-Clause \"Simplified\" License",
		"BSD 2-Clause - Ian Darwin variant",
		"BSD 2-Clause - first lines requirement",
		"BSD-2-Clause Plus Patent License",
		"BSD 2-Clause with views sentence",
		"BSD 3-Clause \"New\" or \"Revised\" License",
		"Apache License 2.0",
		"Mozilla Public License 2.0",
		"ISC License",
	}
	isApproved := false
	for _, approved := range approvedLicenses {
		if license == approved {
			isApproved = true
			break
		}
	}

	moveResult := raidengine.MovementResult{
		Description: "The license for the released software meets the OSI Open Source Definition or the FSF Free Software Definition",
		Function:    utils.CallerPath(0),
		Passed:      isApproved,
		Message:     fmt.Sprintf("Repository uses license: %s", license),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_04
	return moveResult
}
