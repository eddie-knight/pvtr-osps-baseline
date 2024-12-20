package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_03() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "AC_03"
	result = raidengine.StrikeResult{
		Description: "The project's version control system MUST prevent unintentional direct commits against the primary branch.",
		ControlID:   "OSPS-AC-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_03_T01)

	return
}

func AC_03_T01() (moveResult raidengine.MovementResult) {
	protectionData := GetData().Repository.DefaultBranchRef.BranchProtectionRule

	var message string
	if protectionData.RestrictsPushes {
		message = "Branch protection rule restricts pushes"
	} else if protectionData.RequiresApprovingReviews {
		message = "Branch protection rule requires approving reviews"
	}

	moveResult = raidengine.MovementResult{
		Description: "Inspect default branch for a protection rule that restricts pushes",
		Function:    utils.CallerPath(0),
		Passed:      protectionData.RestrictsPushes || protectionData.RequiresApprovingReviews,
		Message:     message,
	}

	return
}
