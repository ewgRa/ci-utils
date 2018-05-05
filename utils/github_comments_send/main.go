package main

import (
	"os"
	"flag"
	"github.com/ewgRa/ci-utils/src/github/comments"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"strings"
	"errors"
	"fmt"
)

func main() {
	file := flag.String("file", "", "Comments, that you want to send to github")
	repo := flag.String("repo", "", "Repository slug")
	pr := flag.Int("pr", 0, "Pull request number")
	flag.Parse()

	if *file == "" || *pr == 0 || *repo == "" {
		flag.Usage()
		os.Exit(1)
	}

	repoPart := strings.Split(*repo, "/")

	if len(repoPart) != 2 {
		panic(errors.New("Can't parse repo"))
	}

	repoOwner := repoPart[0]
	*repo = repoPart[1]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commentsList := comments.ReadComments(*file)

	for _, comment := range commentsList {

		_, _, err := client.PullRequests.CreateComment(ctx, repoOwner, *repo, *pr, comment)

		if err != nil {
			panic(err)
		}
	}
}
