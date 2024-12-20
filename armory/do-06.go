package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_06() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions.",
		ControlID:   "OSPS-DO-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_06_T01)
	return "DO_06", result
}

func DO_06_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_06
	return moveResult
}
