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
		WebCommitSignoffRequired bool
		Name                     string
		HasDiscussionsEnabled    bool
		HasIssuesEnabled         bool
		IsSecurityPolicyEnabled  bool
		DefaultBranchRef         struct {
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
		LicenseInfo struct {
			Name   string
			SpdxId string
			Url    string
		}
		Releases struct {
			TotalCount int
			Nodes      []struct {
				TagName       string
				IsLatest      bool
				ReleaseAssets struct {
					Nodes []struct {
						Name    string
						Url     string
						Content string
					}
				}
			}
		} `graphql:"releases(first: 1, orderBy: {field: CREATED_AT, direction: DESC})"`
		LatestRelease struct {
			Description string
		}
		ContributingGuidelines struct {
			Body         string
			ResourcePath githubv4.URI
		}
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
