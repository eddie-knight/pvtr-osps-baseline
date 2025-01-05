package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_03() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "Any websites, API responses or other services involved in the project development and release MUST be delivered using SSH, HTTPS or other encrypted channels.",
		ControlID:   "OSPS-BR-03",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_03_T01)

	return "BR_03", result
}

func BR_03_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "Reading all URL/URI values from the Security Insights data and checking whether they are prefixed with 'git@', 'https:', or 'ssh:'.",
		Function:    utils.CallerPath(0),
		Passed:      true, // default pass unless a bad link is found
	}

	si := Data.Rest().Insights

	links := []string{
		si.Header.URL,
		si.Header.ProjectSISource,
		si.Project.Homepage,
		si.Project.Roadmap,
		si.Project.Funding,
		si.Project.Documentation.DetailedGuide,
		si.Project.Documentation.CodeOfConduct,
		si.Project.Documentation.QuickstartGuide,
		si.Project.Documentation.ReleaseProcess,
		si.Project.Documentation.SignatureVerification,
		si.Project.Vulnerability.BugBountyProgram,
		si.Project.Vulnerability.SecurityPolicy,
		si.Repository.URL,
		si.Repository.License.URL,
		si.Repository.Security.Assessments.Self.Evidence,
	}

	for _, repo := range si.Project.Repositories {
		links = append(links, repo.URL)
	}

	for _, repo := range si.Repository.Security.Assessments.ThirdParty {
		links = append(links, repo.Evidence)
	}

	for _, tool := range si.Repository.Security.Tools {
		links = append(links, tool.Results.Adhoc.Location)
		links = append(links, tool.Results.CI.Location)
		links = append(links, tool.Results.Release.Location)
	}

	value := make(map[string]string)
	var unknowns int

	for _, url := range links {
		if url == "" {
			continue
		}
		if len(url) <= 3 {
			value[url] = "Unknown, skipped"
			unknowns++
			continue
		}
		if url[:4] == "git@" || url[:6] == "https:" || url[:4] == "ssh:" {
			value[url] = "Passed"
		} else {
			value[url] = "Failed"
			testResult.Passed = false
		}
	}

	testResult.Message = fmt.Sprintf("Secure transport for URLs/URIs: %v", testResult.Passed)
	if unknowns > 0 {
		testResult.Message += fmt.Sprintf(" (!NOTE: %d unknown)", unknowns)
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return testResult
}
