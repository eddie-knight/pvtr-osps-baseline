package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_06() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_06"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All releases MUST provide a descriptive log of functional and security modifications.",
		ControlID:   "OSPS-BR-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_06_T01)

	return
}

func BR_06_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
