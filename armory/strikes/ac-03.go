package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_03() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "AC_03"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s version control system MUST prevent unintentional direct commits against the primary branch.",
		ControlID:   "OSPS-AC-03",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_03_T01)

	return
}

func AC_03_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return
}
