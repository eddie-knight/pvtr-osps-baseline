package armory

import (
	"context"

	"github.com/privateerproj/privateer-sdk/config"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type RepoData struct {
	// Need to update token for this
	// Organization struct {
	// 	Rulesets struct {
	// 		Nodes []struct {
	// 			BypassActors struct {
	// 				TotalCount int
	// 			}
	// 		}
	// 	} `graphql:"rulesets(first: 10)"`
	// } `graphql:"organization(login: $owner)"`

	Repository struct {
		Name                    string
		HasDiscussionsEnabled   bool
		HasIssuesEnabled        bool
		IsSecurityPolicyEnabled bool
		DefaultBranchRef        struct {
			Name          string
			RefUpdateRule struct {
				AllowsDeletions              bool
				AllowsForcePushes            bool
				RequiredApprovingReviewCount bool
			}
		}
		Releases struct {
			TotalCount int
		}
		// BranchProtectionRule	struct{
		// // 	allowsForcePushes			bool
		// // 	requiresApprovingReviews	bool
		// // 	restrictsPushes				bool
		// // 	allowsDeletions				bool
		// 	RequiresStatusChecks		bool
		// }
		LatestRelease struct {
			Description string
		}
		ContributingGuidelines struct {
			Body         string
			ResourcePath githubv4.URI
		}
		BranchProtectionRules struct {
			Nodes []struct {
				AllowsDeletions          bool
				AllowsForcePushes        bool
				RequiresApprovingReviews bool
				RequiresCommitSignatures bool
				RequiresStatusChecks     bool
			}
		} `graphql:"branchProtectionRules(first: 10)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

var GlobalData RepoData

func GetData(c *config.Config) RepoData {
	if GlobalData.Repository.Name != "" {
		return GlobalData
	}

	if c.GetString("token") == "" {
		// TODO: Add unauthenticated data retrieval
		var updatedTactics []string
		for _, tactic := range viper.GetStringSlice("tactics") {
			// append _unauthenticated to each requested tactic name
			updatedTactics = append(updatedTactics, tactic+"_unauthenticated")
		}
		viper.Set("tactics", updatedTactics)
		return GlobalData
	} else {
		return getGraphqlData(c)
	}
}

func getGraphqlData(c *config.Config) RepoData {
	owner := c.GetString("owner")
	repo := c.GetString("repo")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.GetString("token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(repo),
	}

	err := client.Query(context.Background(), &GlobalData, variables)
	if err != nil {
		c.Logger.Error("Error querying GitHub GraphQL API: ", err)
	}
	return GlobalData
}
