package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The version control system MUST require all code contributors to assert that they are legally authorized to commit the associated contributions on every commit.",
		ControlID:   "OSPS-LE-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(LE_01_T01)

	return "LE_01", result
}

// TODO
func LE_01_T01() raidengine.MovementResult {
	orgRequired := Data.GraphQL().Organization.WebCommitSignoffRequired
	repoRequired := Data.GraphQL().Repository.WebCommitSignoffRequired

	required := orgRequired || repoRequired

	moveResult := raidengine.MovementResult{
		Description: "Inspect Org & Repo Policy to Enforce Web SignOff",
		Function:    utils.CallerPath(0),
		Passed:      required,
		Message:     fmt.Sprintf("Web SignOff Enabled: %v", required),
	}

	return moveResult
}
