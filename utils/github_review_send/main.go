package main

import (
	"os"
	"flag"
	"github.com/ewgRa/ci-utils/src/github/comments"
	"github.com/ewgRa/ci-utils/src/github/review"
	"github.com/google/go-github/github"
)

func main() {
	body := flag.String("body", "", "The body text of the pull request review")
	file := flag.String("file", "", "Comments, that you want to send to github")
	repo := flag.String("repo", "", "Repository slug")
	pr := flag.Int("pr", 0, "Pull request number")
	flag.Parse()

	if *file == "" || *pr == 0 || *repo == "" || *body == "" {
		flag.Usage()
		os.Exit(1)
	}

	commentsList := comments.ReadComments(*file)

	event := "REQUEST_CHANGES"

	reviewRequest := &github.PullRequestReviewRequest{
		Body: body,
		Event: &event,
		Comments: commentsList.ToDraftReviewComments(),
	}

	err := review.SendReview(os.Getenv("GITHUB_COMMENTS_SEND_TOKEN"), *repo, *pr, reviewRequest)

	if err != nil {
		panic(err)
	}
}
