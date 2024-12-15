package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_02() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include an explanation of the contribution process.",
		ControlID:   "OSPS-DO-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_02_T01)

	return
}

func DO_02_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_01
	return
}
