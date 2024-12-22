package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_02() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_02_T01)

	return "LE_02", result
}

// TODO
func LE_02_T01() raidengine.MovementResult {
	// TODO Check if this is the correct license data
	approvedLicenses := []string{
		"MIT", "Apache-2.0", "GPL-2.0", "GPL-3.0",
		"LGPL-2.1", "LGPL-3.0", "BSD-2-Clause",
		"BSD-3-Clause", "MPL-2.0", "AGPL-3.0"}

	licenseID := Data.GraphQL().Repository.LicenseInfo.SpdxId
	licenseName := Data.GraphQL().Repository.LicenseInfo.Name

	isApproved := false

	for _, approved := range approvedLicenses {
		if licenseID == approved {
			isApproved = true
			break
		}
	}
	moveResult := raidengine.MovementResult{
		Description: "Verify repository license is OSI/FSF approved",
		Function:    utils.CallerPath(0),
		Passed:      isApproved,
		Message:     fmt.Sprintf("License: %s (%s), Approved: %v", licenseName, licenseID, isApproved),
	}

	return moveResult
}
