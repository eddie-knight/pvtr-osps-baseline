package armory

import (
	"context"
	"log"

	"github.com/privateerproj/privateer-sdk/config"
	"github.com/shurcooL/githubv4"
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
		Releases                struct {
			TotalCount int
		}
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

	owner := c.GetString("owner")
	repo := c.GetString("repo")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.GetString("token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	return makeQuery(client, owner, repo)
}

func makeQuery(client *githubv4.Client, owner, name string) RepoData {
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	err := client.Query(context.Background(), &GlobalData, variables)
	if err != nil {
		log.Print(err)
	}

	return GlobalData
}
