package main

import (
	"os"
	"flag"
	"github.com/ewgRa/ci-utils/src/github/comments"
)

// Send comments to pull request under personal token
func main() {
	file := flag.String("file", "", "Comments, that you want to send to github")
	repo := flag.String("repo", "", "Repository slug")
	pr := flag.Int("pr", 0, "Pull request number")
	flag.Parse()

	if *file == "" || *pr == 0 || *repo == "" {
		flag.Usage()
		os.Exit(1)
	}

	commentsList := comments.ReadComments(*file)

	err := comments.SendComments(os.Getenv("GITHUB_COMMENTS_SEND_TOKEN"), *repo, *pr, commentsList)

	if err != nil {
		panic(err)
	}
}
