package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_05() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_05"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include a mechanism for reporting defects.",
		ControlID:   "OSPS-DO-05",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_05_T01)

	return
}

func DO_05_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_01
	return
}
