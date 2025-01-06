package armory

import (
	"github.com/ossf/si-tooling/v2/si"
	"github.com/privateerproj/privateer-sdk/pluginkit"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_07() (string, pluginkit.TestSetResult) {
	result := pluginkit.TestSetResult{
		Description: "All released software assets MUST be signed or accounted for in a signed manifest including each asset’s cryptographic hashes.",
		ControlID:   "OSPS-BR-07",
		Tests:       make(map[string]pluginkit.TestResult),
	}

	result.ExecuteTest(BR_07_T01)
	// TODO: Ensure the identified attestation exists

	return "BR_07", result
}

func BR_07_T01() pluginkit.TestResult {
	testResult := pluginkit.TestResult{
		Description: "Check Security Insights Release Attestations for SLSA predicate-uri.",
		Function:    utils.CallerPath(0),
	}

	attestations := Data.Rest().Insights.Repository.Release.Attestations

	var value []si.Attestation

	for _, attestation := range attestations {
		if attestation.PredicateURI == "https://slsa.dev/provenance/v1" {
			testResult.Passed = true
			testResult.Message = "Attestsation with SLSA predicate type."
			testResult.Value = attestation
			value = append(value, attestation)
		}
	}
	testResult.Value = value
	return testResult
}
