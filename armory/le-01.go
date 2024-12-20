package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The version control system MUST require all code contributors to assert that they are legally authorized to commit the associated contributions on every commit.",
		ControlID:   "OSPS-LE-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_01_T01)

	return "LE_01", result
}

// TODO
func LE_01_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_01
	return moveResult
}
