package main

import (
	"os"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ewgRa/ci-utils/src/github/comments"
)

// Comments diff parse comments that you want to send and intersect it with exists comments to avoid double commenting
func main() {
	commentsFile := flag.String("comments", "", "Comments, that you want to send to github")
	existsCommentsFile := flag.String("exists-comments", "", "Exists comments from github")
	flag.Parse()

	if *commentsFile == "" || *existsCommentsFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	existsComments := comments.ReadComments(*existsCommentsFile)
	commentsList := comments.ReadComments(*commentsFile)

	var comments comments.GithubComments

	for _, comment := range commentsList {
		if !existsComments.Has(comment) {
			comments = append(comments, comment)
		}
	}

	if len(comments) > 0 {
		json, err := json.Marshal(comments)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(json))
	}
}
