package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_07() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST provide design documentation demonstrating all actions and actors within the system.",
		ControlID:   "OSPS-AC-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_07_T01)

	return "DO_07", result
}

// TODO
func DO_07_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_07
	return moveResult
}
