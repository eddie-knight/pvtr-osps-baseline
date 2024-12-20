package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_03() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "All released software assets MUST be delivered with a machine-readable list of all direct and transitive internal software dependencies with their associated version identifiers.",
		ControlID:   "OSPS-QA-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_03_T01)

	return "QA_03", result
}

// TODO
func QA_03_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to QA_03
	return moveResult
}
