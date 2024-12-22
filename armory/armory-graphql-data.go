package armory

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GraphqlData struct {
	// Need to update token for this
	Organization struct {
		RequiresTwoFactorAuthentication bool
		WebCommitSignoffRequired        bool
	} `graphql:"organization(login: $owner)"`

	Repository struct {
		Name                     string
		HasDiscussionsEnabled    bool
		HasIssuesEnabled         bool
		IsSecurityPolicyEnabled  bool
		WebCommitSignoffRequired bool
		LicenseInfo              struct {
			SpdxId string
			Name   string
			Url    string
		}
		DefaultBranchRef struct {
			Name          string
			RefUpdateRule struct { // Docs say this works for non-admin viewers, but I haven't managed to do that yet
				AllowsDeletions              bool
				AllowsForcePushes            bool
				RequiredApprovingReviewCount int
			}
			BranchProtectionRule struct {
				RestrictsPushes          bool // This didn't give an accurate result
				RequiresApprovingReviews bool // This gave an accurate result
				RequiresCommitSignatures bool
				RequiresStatusChecks     bool
			}
		}
		Releases struct {
			TotalCount int
		}
		LatestRelease struct {
			Description string
		}
		ContributingGuidelines struct {
			Body         string
			ResourcePath githubv4.URI
		}
		// Collaborators struct {
		// 	edges struct {
		// 		node struct {
		// 			Cursor            string
		// 			PermissionSources struct {
		// 				Permission string
		// 			}
		// 			node struct {
		// 				Login string
		// 			}
		// 		}
		// 	}
		// }
	} `graphql:"repository(owner: $owner, name: $name)"`
}

func (r *ArmoryData) loadGraphQLData() error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GlobalConfig.GetString("token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner": githubv4.String(GlobalConfig.GetString("owner")),
		"name":  githubv4.String(GlobalConfig.GetString("repo")),
	}

	err := client.Query(context.Background(), &Data.graphql, variables)
	if err != nil {
		Logger.Error(fmt.Sprintf("Error querying GitHub GraphQL API: %s", err.Error()))
	}
	return nil
}
