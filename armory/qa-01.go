package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func QA_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project's source code MUST be publicly readable and have a static URL.",
		ControlID:   "OSPS-QA-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(QA_01_T01)

	return "QA_01", result
}

func QA_01_T01() raidengine.MovementResult {
	gotRepoData := Data.Rest().Repo.Name != ""
	isPrivate := Data.Rest().Repo.Private

	moveResult := raidengine.MovementResult{
		Description: "Verifying that the GitHub repository is public at the target URL.",
		Function:    utils.CallerPath(0),
		Passed:      gotRepoData && !isPrivate,
	}

	if !gotRepoData {
		moveResult.Message = "Repository data not found"
	} else {
		moveResult.Message = fmt.Sprintf("Public Repo: %t", !isPrivate)
	}

	return moveResult
}
