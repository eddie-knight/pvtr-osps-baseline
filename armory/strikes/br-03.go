package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_03() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "Any websites, API responses or other services involved in the project development and release MUST be delivered using SSH, HTTPS or other encrypted channels.",
		ControlID:   "OSPS-BR-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_03_T01)

	return
}

func BR_03_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
