package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_05() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "AC_05"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated.",
		ControlID:   "OSPS-AC-05",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_05_T01)

	return
}

func AC_05_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return
}
