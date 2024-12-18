package armory

import (
	"strings"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_02() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include an explanation of the contribution process.",
		ControlID:   "OSPS-DO-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_02_T01)

	return
}

func DO_02_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "Inspecting the GitHub repo's recommended contributing guidelines to ensure it is not empty.",
		Function:    utils.CallerPath(0),
	}

	data := GetData(Config)
	body := data.Repository.ContributingGuidelines.Body
	if strings.Contains(body, "Contributing") {
		moveResult.Passed = true
	} else {
		moveResult.Passed = false
	}

	return
}
