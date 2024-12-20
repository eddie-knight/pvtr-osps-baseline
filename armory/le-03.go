package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_03() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The license for the source code MUST be maintained in a standard location within the project’s repository.",
		ControlID:   "OSPS-LE-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_03_T01)

	return "LE_03", result
}

// TODO
func LE_03_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_03
	return moveResult
}
