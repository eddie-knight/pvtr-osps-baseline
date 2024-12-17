package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_04() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "LE_04"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_04_T01)

	return
}

// TODO
func LE_04_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_04
	return
}
