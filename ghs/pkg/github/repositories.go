package github

import (
	"context"

	"github.com/google/go-github/v30/github"
)

func (gh *GitHub) GetRepositories(owner string) ([]*github.Repository, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	repos, _, err := gh.client.Repositories.List(context.Background(), owner, opt)
	return repos, err
}
