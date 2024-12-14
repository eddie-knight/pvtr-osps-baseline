package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_02() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "LE_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		ControlID:   "OSPS-LE-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_02_T01)

	return
}

// TODO
func LE_02_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to LE_02
	return
}
