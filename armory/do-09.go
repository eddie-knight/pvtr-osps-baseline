package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_09() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_09"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include descriptions of all external input and output interfaces of the released software assets.",
		ControlID:   "OSPS-DO-09",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_09_T01)

	return
}

// TODO
func DO_09_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_09
	return
}
