package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_06() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project’s version control system MUST require multi-factor authentication that does not include SMS for users when modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_06_T01)

	return "AC_06", result
}

func AC_06_T01() raidengine.MovementResult {
	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return moveResult
}
