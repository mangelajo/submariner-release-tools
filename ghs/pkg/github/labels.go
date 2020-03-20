package github

import (
	"context"

	"github.com/google/go-github/v30/github"
)

func (gh *GitHub) CreateLabel(owner, repo, name, color, description string) (*github.Label, *github.Response, error) {
	item := github.Label{
		Name:        &name,
		Color:       &color,
		Description: &description}

	return gh.client.Issues.CreateLabel(context.Background(), owner, repo, &item)
}
