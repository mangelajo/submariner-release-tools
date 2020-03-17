package github

import (
	"context"
	"fmt"
	"os"

	goGithub "github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

func GetGithubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

type GitHub struct {
	client *goGithub.Client
}

func NewGitHub() (*GitHub, error) {
	ctx := context.Background()
	token := GetGithubToken()
	if token == "" {
		return nil, fmt.Errorf("No GITHUB_TOKEN provided in environment")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &GitHub{client: goGithub.NewClient(tc)}, nil
}

func (gh *GitHub) GetRepositories(owner string) ([]*goGithub.Repository, error) {
	opt := &goGithub.RepositoryListOptions{
		ListOptions: goGithub.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	repos, _, err := gh.client.Repositories.List(context.Background(), owner, opt)
	return repos, err
}

func (gh *GitHub) GetMilestones(owner, repo string) ([]*goGithub.Milestone, *goGithub.Response, error) {
	opt := &goGithub.MilestoneListOptions{
		ListOptions: goGithub.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	return gh.client.Issues.ListMilestones(context.Background(), owner, repo, opt)

}

func (gh *GitHub) CreateMilestone(owner, repo, milestone string) (*goGithub.Milestone, *goGithub.Response, error) {
	item := goGithub.Milestone{Title: &milestone}
	return gh.client.Issues.CreateMilestone(context.Background(), owner, repo, &item)
}

func (gh *GitHub) CreateLabel(owner, repo, name, color, description string) (*goGithub.Label, *goGithub.Response, error) {
	item := goGithub.Label{
		Name:        &name,
		Color:       &color,
		Description: &description}

	return gh.client.Issues.CreateLabel(context.Background(), owner, repo, &item)
}
