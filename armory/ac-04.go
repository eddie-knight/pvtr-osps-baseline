package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_04() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "AC_04"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s version control system MUST prevent unintentional deletion of the primary branch.",
		ControlID:   "OSPS-AC-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_04_T01)

	return
}

func AC_04_T01() (moveResult raidengine.MovementResult) {
	allowed := GetData().Repository.DefaultBranchRef.RefUpdateRule.AllowsDeletions
	branchName := GetData().Repository.DefaultBranchRef.Name
	msg := fmt.Sprintf("Branch Protection Prevents Deletion: %v", !allowed)
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Passed:      !allowed,
		Value:       branchName,
		Message:     msg,
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return
}
