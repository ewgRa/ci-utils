package main

import (
	"os"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ewgRa/ci-utils/github/comments"
)

// DiffLiner parse github diff content, taken by github api call (e.g. curl -H "Accept: application/vnd.github.v3.diff.json" https://api.github.com/repos/ru-de/faq/pulls/377)
// As output it respond with json data, that says which one position in diff belongs to changed line in file.
// This position needed for call comment API endpoint - https://developer.github.com/v3/pulls/comments/#create-a-comment
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

	for _, comment := range commentsList {
		if !existsComments.Has(comment) {
			json, err := json.Marshal(comment)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(json))
		}
	}
}
