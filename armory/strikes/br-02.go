package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_02() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All releases and released software assets MUST be assigned a unique version identifier for each release intended to be used by users.",
		ControlID:   "OSPS-BR-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_02_T01)

	return
}

func BR_02_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
