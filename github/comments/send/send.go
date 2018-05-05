package main

import (
	"os"
	"flag"
	"github.com/ewgRa/ci-utils/github/comments"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
)

func main() {
	file := flag.String("file", "", "Comments, that you want to send to github")
	flag.Parse()

	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commentsList := comments.ReadComments(*file)

	for _, comment := range commentsList {

		_, _, err := client.PullRequests.CreateComment(ctx, "ru-de", "faq", 377, comment)

		if err != nil {
			panic(err)
		}
	}
}
