package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_07() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_07"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST provide design documentation demonstrating all actions and actors within the system.",
		ControlID:   "OSPS-AC-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_07_T01)

	return
}

// TODO
func DO_07_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_07
	return
}
