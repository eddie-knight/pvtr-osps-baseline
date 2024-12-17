package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_06() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_06"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All releases MUST provide a descriptive log of functional and security modifications.",
		ControlID:   "OSPS-DO-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_06_T01)

	return
}

func DO_06_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_01
	return
}
