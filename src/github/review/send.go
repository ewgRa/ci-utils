package review

import (
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"strings"
	"errors"
	"context"
	"fmt"
	"io/ioutil"
)

func SendReview(accessToken, repo string, prNumber int, reviewRequest *github.PullRequestReviewRequest) error {
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

	r, response, err := client.PullRequests.CreateReview(ctx, repoOwner, repo, prNumber, reviewRequest)

	if err != nil {
		b, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(b),r, reviewRequest)
		return err
	}

	return nil
}
