package github

import (
	"context"

	"github.com/google/go-github/v30/github"
)

func (gh *GitHub) GetMilestones(owner, repo string) ([]*github.Milestone, *github.Response, error) {
	opt := &github.MilestoneListOptions{
		ListOptions: github.ListOptions{PerPage: 100}, //TODO: Support more than the max 100 per-page repos
	}
	return gh.client.Issues.ListMilestones(context.Background(), owner, repo, opt)

}

func (gh *GitHub) CreateMilestone(owner, repo, milestone string) (*github.Milestone, *github.Response, error) {
	item := github.Milestone{Title: &milestone}
	return gh.client.Issues.CreateMilestone(context.Background(), owner, repo, &item)
}
