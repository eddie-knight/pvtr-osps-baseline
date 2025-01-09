package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"

	"github.com/privateerproj/privateer-sdk/utils"
)

func LE_01() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "The version control system MUST require all code contributors to assert that they are legally authorized to commit the associated contributions on every commit.",
		ControlID:   "OSPS-LE-01",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(LE_01_T01)

	return "LE_01", result
}

func LE_01_T01() pluginkit.TestResult {
	orgRequired := Data.GraphQL().Organization.WebCommitSignoffRequired
	repoRequired := Data.GraphQL().Repository.WebCommitSignoffRequired

	required := orgRequired || repoRequired

	moveResult := pluginkit.TestResult{
		Description: "Inspect Org & Repo Policy to Enforce Web SignOff",
		Function:    utils.CallerPath(0),
		Passed:      required,
		Message:     fmt.Sprintf("Web SignOff Enabled: %v", required),
	}

	return moveResult
}
