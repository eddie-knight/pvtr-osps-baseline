package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_09() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST include descriptions of all external input and output interfaces of the released software assets.",
		ControlID:   "OSPS-DO-09",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_09_T01)

	return "DO_09", result
}

// TODO
func DO_09_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_09
	return moveResult
}
