package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_01() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles.",
		ControlID:   "OSPS-DO-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_01_T01)

	return
}

// TODO
func DO_01_T01() raidengine.MovementResult {
	data := GetData(Config)

	return raidengine.MovementResult{
		Description: "Inspecting whether issues or discussions are enabled on the repo.",
		Function:    utils.CallerPath(0),
		Passed:      data.Repository.HasDiscussionsEnabled, // TODO: || data.Repository.HasIssuesEnabled
	}
}
