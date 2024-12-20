package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project’s build and release pipelines MUST NOT execute arbitrary code that is input from outside of the build script.",
		ControlID:   "OSPS-BR-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_01_T01)

	return "BR_01", result
}

// TODO
func BR_01_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}
