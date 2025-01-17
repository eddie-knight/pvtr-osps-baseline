package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_02() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All releases and released software assets MUST be assigned a unique version identifier for each release intended to be used by users.",
		ControlID:   "OSPS-BR-02",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_02_T01)
	if result.Tests["BR_02_T01"].Value.(int) > 0 {
		result.ExecuteTest(BR_02_T02)
	}
	return "BR_02", result
}

func BR_02_T01() pluginkit.TestResult {
	return countReleases()
}

func BR_02_T02() pluginkit.TestResult {
	releases := Data.Rest().Repo.Releases

	releaseNames := make(map[string]int)
	var errorCount int
	for _, release := range releases {
		if release.Name == "" {
			errorCount++
			GlobalConfig.Logger.Error("Release %v has no name!", release.Name)
		} else if _, ok := releaseNames[release.Name]; ok {
			errorCount++
			GlobalConfig.Logger.Error(fmt.Sprintf(
				"Release id: %v has the same name as another release: %s", release.Id, release.Name))
		} else {
			releaseNames[release.Name] = release.Id
		}
	}

	testResult := pluginkit.TestResult{
		Description: "Ensure all releases have a unique name",
		Function:    utils.CallerPath(0),
		Passed:      errorCount == 0,
		Message:     fmt.Sprintf("Non-unique release names: %v", errorCount),
	}

	return testResult
}
