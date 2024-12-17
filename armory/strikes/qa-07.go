package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_07() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "QA_07"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All proposed changes to the project’s codebase must be automatically evaluated against a documented policy for known vulnerabilities and blocked in the event of violations except when declared and supressed as non exploitable.",
		ControlID:   "OSPS-QA-07",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_07_T01)

	return
}

// TODO
func QA_07_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_07
	return
}
