package comments

import (
	"encoding/json"
	"github.com/google/go-github/github"
	"io/ioutil"
)

type GithubComments []*github.PullRequestComment

func (cs GithubComments) Has(comment *github.PullRequestComment) bool {
	for _, cmt := range cs {
		if cmt.GetBody() == comment.GetBody() && cmt.GetPath() == comment.GetPath() && cmt.GetPosition() == comment.GetPosition() {
			return true
		}
	}

	return false
}

func ReadComments(file string) GithubComments {
	var commentsList GithubComments

	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	if len(content) == 0 {
		return commentsList
	}

	err = json.Unmarshal(content, &commentsList)

	if err != nil {
		panic(err)

	}

	return commentsList
}
