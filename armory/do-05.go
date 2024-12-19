package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_05() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_05"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include a mechanism for reporting defects.",
		ControlID:   "OSPS-DO-05",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_05_T01)

	return
}

func DO_05_T01() (moveResult raidengine.MovementResult) {
	enabled := GetData().Repository.HasIssuesEnabled
	moveResult = raidengine.MovementResult{
		Description: "Checking to see whether the GitHub repo has issues enabled",
		Function:    utils.CallerPath(0),
		Passed:      enabled,
		Message:     fmt.Sprintf("Issues Enabled: %v", enabled),
	}

	return
}
