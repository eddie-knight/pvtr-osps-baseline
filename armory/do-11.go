package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_11() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_11"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources.",
		ControlID:   "OSPS-DO-11",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_11_T01)

	return
}

// TODO
func DO_11_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_11
	return
}
