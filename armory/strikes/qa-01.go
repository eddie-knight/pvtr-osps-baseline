package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_01() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "QA_01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s source code MUST be publicly readable and have a static URL.",
		ControlID:   "OSPS-QA-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_01_T01)

	return
}

// TODO
func QA_01_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_01
	return
}
