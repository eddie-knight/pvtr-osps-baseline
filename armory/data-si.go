package armory

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type SecurityInsights struct {
	rawData    []byte
	Header     Header     `yaml:"header"`
	Project    Project    `yaml:"project"`
	Repository Repository `yaml:"repository"`
}

type Header struct {
	LastReviewed    string  `yaml:"last-reviewed"`
	LastUpdated     string  `yaml:"last-updated"`
	SchemaVersion   string  `yaml:"schema-version"`
	URL             string  `yaml:"url"`
	Comment         *string `yaml:"comment"`
	ProjectSISource *string `yaml:"project-si-source"`
}

type Assessment struct {
	Comment  string `yaml:"comment"`
	Name     string `yaml:"name"`
	Evidence string `yaml:"evidence"`
	Date     string `yaml:"date"`
}

type Attestation struct {
	Name         string `yaml:"name"`
	Location     string `yaml:"location"`
	PredicateURI string `yaml:"predicate-uri"`
	Comment      string `yaml:"comment"`
}

type Contact struct {
	Name        string `yaml:"name"`
	Primary     bool   `yaml:"primary"`
	Affiliation string `yaml:"affiliation"`
	Email       string `yaml:"email"`
	Social      string `yaml:"social"`
}

type License struct {
	URL        string `yaml:"url"`
	Expression string `yaml:"expression"`
}

type Link struct {
	URI     string `yaml:"uri"`
	Comment string `yaml:"comment"`
}

type Project struct {
	Name           string     `yaml:"name"`
	Homepage       string     `yaml:"homepage"`
	Roadmap        string     `yaml:"roadmap"`
	Funding        string     `yaml:"funding"`
	Administrators []Contact  `yaml:"administrators"`
	Repositories   []Repo     `yaml:"repositories"`
	Vulnerability  VulnReport `yaml:"vulnerability-reporting"`
	Documentation  Docs       `yaml:"documentation"`
}

type Repo struct {
	Name    string `yaml:"name"`
	Comment string `yaml:"comment"`
	URL     string `yaml:"url"`
}

type VulnReport struct {
	ReportsAccepted    bool     `yaml:"reports-accepted"`
	BugBountyAvailable bool     `yaml:"bug-bounty-available"`
	BugBountyProgram   string   `yaml:"bug-bounty-program"`
	Contact            Contact  `yaml:"contact"`
	Comment            string   `yaml:"comment"`
	SecurityPolicy     string   `yaml:"security-policy"`
	PGPKey             string   `yaml:"pgp-key"`
	InScope            []string `yaml:"in-scope"`
	OutOfScope         []string `yaml:"out-of-scope"`
}

type Docs struct {
	DetailedGuide         string `yaml:"detailed-guide"`
	CodeOfConduct         string `yaml:"code-of-conduct"`
	QuickstartGuide       string `yaml:"quickstart-guide"`
	ReleaseProcess        string `yaml:"release-process"`
	SignatureVerification string `yaml:"signature-verification"`
}

type Repository struct {
	Status                        string       `yaml:"status"`
	URL                           string       `yaml:"url"`
	AcceptsChangeRequest          bool         `yaml:"accepts-change-request"`
	AcceptsAutomatedChangeRequest bool         `yaml:"accepts-automated-change-request"`
	BugFixesOnly                  bool         `yaml:"bug-fixes-only"`
	NoThirdPartyPackages          bool         `yaml:"no-third-party-packages"`
	CoreTeam                      []Contact    `yaml:"core-team"`
	License                       License      `yaml:"license"`
	Security                      SecurityInfo `yaml:"security"`
	Documentation                 Docs         `yaml:"documentation"`
	Release                       Release      `yaml:"release"`
}

type SecurityInfo struct {
	Assessments Assessments `yaml:"assessments"`
	Champions   []Contact   `yaml:"champions"`
	Tools       []Tool      `yaml:"tools"`
}

type Assessments struct {
	Self       Assessment   `yaml:"self"`
	ThirdParty []Assessment `yaml:"third-party"`
}

type Tool struct {
	Name        string      `yaml:"name"`
	Type        string      `yaml:"type"`
	Version     string      `yaml:"version"`
	Comment     string      `yaml:"comment"`
	Rulesets    []string    `yaml:"rulesets"`
	Integration Integration `yaml:"integration"`
	Results     Results     `yaml:"results"`
}

type Integration struct {
	Adhoc   bool `yaml:"adhoc"`
	CI      bool `yaml:"ci"`
	Release bool `yaml:"release"`
}

type Results struct {
	Adhoc   Attestation `yaml:"adhoc"`
	CI      Attestation `yaml:"ci"`
	Release Attestation `yaml:"release"`
}

type Release struct {
	AutomatedPipeline  bool          `yaml:"automated-pipeline"`
	DistributionPoints []Link        `yaml:"distribution-points"`
	Changelog          string        `yaml:"changelog"`
	License            License       `yaml:"license"`
	Attestations       []Attestation `yaml:"attestations"`
}

type SIHeader struct {
	SchemaVersion string `yaml:"schema-version"`
	ChangeLogURL  string `yaml:"changelog"`
	LicenseURL    string `yaml:"license"`
}

func (s *SecurityInsights) Ingest() {
	// TODO: support multiple schema versions
	major, minor, patch := ParseVersion(s.Header.SchemaVersion)
	if major != 2 {
		Logger.Error(fmt.Sprintf("the requested Security Insights schema version is not currently supported: \n%s", s.Header.SchemaVersion))
		return
	}
	if minor != 0 || patch != 0 {
		Logger.Warn("the provided Security Insights schema version may not be supported: '%s' (known supported is v2.0.0)", s.Header.SchemaVersion)
	}
	err := yaml.Unmarshal(s.rawData, s)
	if err != nil {
		Logger.Error(fmt.Sprintf("error unmarshalling security insights byte data into struct: %s", err.Error()))
		return
	}
}

func ParseVersion(version string) (int, int, int) {
	splitVersion := strings.Split(version, ".")
	if len(splitVersion) == 3 {
		major, _ := strconv.Atoi(splitVersion[0])
		minor, _ := strconv.Atoi(splitVersion[1])
		patch, _ := strconv.Atoi(splitVersion[2])
		return major, minor, patch
	}
	if len(splitVersion) == 2 {
		major, _ := strconv.Atoi(splitVersion[0])
		minor, _ := strconv.Atoi(splitVersion[1])
		return major, minor, 0
	}
	if len(splitVersion) == 1 {
		major, _ := strconv.Atoi(splitVersion[0])
		return major, 0, 0
	}
	Logger.Error("found invalid Security Insights schema version: %s", version)
	return 0, 0, 0
}
