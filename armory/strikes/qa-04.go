package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_04() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "QA_04"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Any automated status checks for commits MUST pass or require manual acknowledgement prior to merge.",
		ControlID:   "OSPS-QA-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_04_T01)

	return
}

func QA_04_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_04
	return
}
