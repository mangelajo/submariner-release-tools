package github

import (
	"context"

	"github.com/google/go-github/v30/github"
)

func (gh *GitHub) GetIssue(owner, repo string, id int) (*github.Issue, *github.Response, error) {
	return gh.client.Issues.Get(context.Background(), owner, repo, id)
}
