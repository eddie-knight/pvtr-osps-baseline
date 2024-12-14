package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_12() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_12"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST contain instructions to verify the integrity and authenticity of the release assets, including the expected identity of the person or process authoring the software release.",
		ControlID:   "OSPS-DO-12",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_12_T01)

	return
}

// TODO
func DO_12_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_12
	return
}
