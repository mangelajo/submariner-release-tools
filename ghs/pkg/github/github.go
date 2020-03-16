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

	return &GitHub{ client: goGithub.NewClient(tc)} , nil
}

func (gh *GitHub) GetRepositories(owner string) ([] *goGithub.Repository, error) {
	opt := &goGithub.RepositoryListOptions{
		ListOptions: goGithub.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	repos, _, err := gh.client.Repositories.List(context.Background(), owner, opt)
	return repos, err
}

func (gh *GitHub) GetMilestones(owner, repo string) ([]*goGithub.Milestone, error) {
	opt := &goGithub.MilestoneListOptions{
		ListOptions: goGithub.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	milestones, _, err := gh.client.Issues.ListMilestones(context.Background(), owner, repo, opt)

	return milestones, err
}

func (gh *GitHub) CreateMilestone(owner, repo, milestone string) (*goGithub.Milestone, error) {
	item := goGithub.Milestone{Title: &milestone}
	milestoneObj, _, err := gh.client.Issues.CreateMilestone(context.Background(), owner, repo, &item)

	return milestoneObj, err
}