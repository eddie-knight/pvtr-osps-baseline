package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_06() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The version control system MUST NOT contain generated executable artifacts.",
		ControlID:   "OSPS-QA-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_06_T01)

	return "QA_06", result
}

func QA_06_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_06
	return moveResult
}
