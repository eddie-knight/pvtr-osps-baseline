package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_07() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_07"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All released software assets MUST be signed or accounted for in a signed manifest including each asset’s cryptographic hashes.",
		ControlID:   "OSPS-BR-07",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_07_T01)

	return
}

func BR_07_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
