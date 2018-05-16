package comments

import (
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"strings"
	"errors"
	"context"
)

func SendComments(accessToken, repo string, prNumber int, commentsList GithubComments) error {
	repoPart := strings.Split(repo, "/")

	if len(repoPart) != 2 {
		return errors.New("Can't parse repo")
	}

	repoOwner := repoPart[0]
	repo = repoPart[1]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for _, comment := range commentsList {

		_, _, err := client.PullRequests.CreateComment(ctx, repoOwner, repo, prNumber, comment)

		if err != nil {
			return err
		}
	}

	return nil
}
