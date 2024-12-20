package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_02() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project’s version control system MUST restrict collaborator permissions to the lowest available privileges by default.",
		ControlID:   "OSPS-AC-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_02_T01)

	return "AC_02", result
}

func AC_02_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return moveResult
}
