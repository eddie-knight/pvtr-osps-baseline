package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_02() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_02_T01)

	return "LE_02", result
}

// TODO
func LE_02_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_02
	return moveResult
}
