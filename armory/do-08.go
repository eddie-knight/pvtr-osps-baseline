package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_08() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST include a policy that defines a threshold for remediation of SCA findings related to vulnerabilities and licenses.",
		ControlID:   "OSPS-DO-08",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_08_T01)

	return "DO_08", result
}

// TODO
func DO_08_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to DO_08
	return moveResult
}
