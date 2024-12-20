package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_03() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST provide user guides for all basic functionality.",
		ControlID:   "OSPS-DO-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_03_T01)

	return "DO_03", result
}

func DO_03_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: This needs security insights data
	return moveResult
}
