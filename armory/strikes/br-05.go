package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_05() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_05"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All build and release pipelines MUST use standardized tooling where available to ingest dependencies at build time.",
		ControlID:   "OSPS-BR-05",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_05_T01)

	return
}

func BR_05_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
