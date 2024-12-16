package armory

import (
	"time"

	"github.com/privateerproj/privateer-sdk/config"
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	Config *config.Config
	Armory = raidengine.Armory{
		Tactics: map[string][]raidengine.Strike{
			"dev": {
				DO_01,
			},
			"maturity_1": {
				AC_01,
				AC_02,
				AC_03,
				AC_04,
				BR_01,
				BR_02,
				BR_03,
				DO_01,
				DO_02,
				LE_01,
				LE_02,
				LE_03,
				QA_01,
				QA_02,
			},
			"maturity_2": {
				AC_05,
				BR_04,
				BR_05,
				BR_06,
				BR_07,
				DO_03,
				DO_04,
				DO_05,
				DO_06,
				DO_07,
				DO_11,
				DO_12,
				LE_04,
				QA_03,
				QA_04,
				QA_05,
				QA_06,
			},
			"maturity_3": {
				AC_06,
				DO_08,
				DO_09,
				DO_10,
				QA_07,
			},
		},
	}
)

// Owner represents the owner of the repository.
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// License represents the license information for the repository.
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SPDXID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

// Organization represents the organization information for the repository.
type Organization struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// RepoData represents the main repository object.
type RepoData struct {
	ID                  int                    `json:"id"`
	NodeID              string                 `json:"node_id"`
	Name                string                 `json:"name"`
	FullName            string                 `json:"full_name"`
	Private             bool                   `json:"private"`
	Owner               Owner                  `json:"owner"`
	HTMLURL             string                 `json:"html_url"`
	Description         string                 `json:"description"`
	Fork                bool                   `json:"fork"`
	URL                 string                 `json:"url"`
	ForksURL            string                 `json:"forks_url"`
	KeysURL             string                 `json:"keys_url"`
	CollaboratorsURL    string                 `json:"collaborators_url"`
	TeamsURL            string                 `json:"teams_url"`
	HooksURL            string                 `json:"hooks_url"`
	IssueEventsURL      string                 `json:"issue_events_url"`
	EventsURL           string                 `json:"events_url"`
	AssigneesURL        string                 `json:"assignees_url"`
	BranchesURL         string                 `json:"branches_url"`
	TagsURL             string                 `json:"tags_url"`
	BlobsURL            string                 `json:"blobs_url"`
	GitTagsURL          string                 `json:"git_tags_url"`
	GitRefsURL          string                 `json:"git_refs_url"`
	TreesURL            string                 `json:"trees_url"`
	StatusesURL         string                 `json:"statuses_url"`
	LanguagesURL        string                 `json:"languages_url"`
	StargazersURL       string                 `json:"stargazers_url"`
	ContributorsURL     string                 `json:"contributors_url"`
	SubscribersURL      string                 `json:"subscribers_url"`
	SubscriptionURL     string                 `json:"subscription_url"`
	CommitsURL          string                 `json:"commits_url"`
	GitCommitsURL       string                 `json:"git_commits_url"`
	CommentsURL         string                 `json:"comments_url"`
	IssueCommentURL     string                 `json:"issue_comment_url"`
	ContentsURL         string                 `json:"contents_url"`
	CompareURL          string                 `json:"compare_url"`
	MergesURL           string                 `json:"merges_url"`
	ArchiveURL          string                 `json:"archive_url"`
	DownloadsURL        string                 `json:"downloads_url"`
	IssuesURL           string                 `json:"issues_url"`
	PullsURL            string                 `json:"pulls_url"`
	MilestonesURL       string                 `json:"milestones_url"`
	NotificationsURL    string                 `json:"notifications_url"`
	LabelsURL           string                 `json:"labels_url"`
	ReleasesURL         string                 `json:"releases_url"`
	DeploymentsURL      string                 `json:"deployments_url"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	PushedAt            time.Time              `json:"pushed_at"`
	GitURL              string                 `json:"git_url"`
	SSHURL              string                 `json:"ssh_url"`
	CloneURL            string                 `json:"clone_url"`
	SVNURL              string                 `json:"svn_url"`
	Homepage            string                 `json:"homepage"`
	Size                int                    `json:"size"`
	StargazersCount     int                    `json:"stargazers_count"`
	WatchersCount       int                    `json:"watchers_count"`
	Language            string                 `json:"language"`
	HasIssues           bool                   `json:"has_issues"`
	HasProjects         bool                   `json:"has_projects"`
	HasDownloads        bool                   `json:"has_downloads"`
	HasWiki             bool                   `json:"has_wiki"`
	HasPages            bool                   `json:"has_pages"`
	HasDiscussions      bool                   `json:"has_discussions"`
	ForksCount          int                    `json:"forks_count"`
	MirrorURL           *string                `json:"mirror_url"`
	Archived            bool                   `json:"archived"`
	Disabled            bool                   `json:"disabled"`
	OpenIssuesCount     int                    `json:"open_issues_count"`
	License             *License               `json:"license"`
	AllowForking        bool                   `json:"allow_forking"`
	IsTemplate          bool                   `json:"is_template"`
	WebCommitSignoffReq bool                   `json:"web_commit_signoff_required"`
	Topics              []string               `json:"topics"`
	Visibility          string                 `json:"visibility"`
	Forks               int                    `json:"forks"`
	OpenIssues          int                    `json:"open_issues"`
	Watchers            int                    `json:"watchers"`
	DefaultBranch       string                 `json:"default_branch"`
	TempCloneToken      *string                `json:"temp_clone_token"`
	CustomProperties    map[string]interface{} `json:"custom_properties"`
	Organization        *Organization          `json:"organization"`
	NetworkCount        int                    `json:"network_count"`
	SubscribersCount    int                    `json:"subscribers_count"`
}
