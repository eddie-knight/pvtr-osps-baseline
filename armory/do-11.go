package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_11() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources.",
		ControlID:   "OSPS-DO-11",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_11_T01)

	return "DO_11", result
}

// TODO
func DO_11_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_11
	return moveResult
}
