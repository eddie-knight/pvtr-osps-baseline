package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_10() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST include a policy to address SCA violations prior to any release.",
		ControlID:   "OSPS-DO-10",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_10_T01)

	return "DO_10", result
}

// TODO
func DO_10_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_10
	return moveResult
}
