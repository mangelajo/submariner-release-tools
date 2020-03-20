package github

import (
	"context"
	"fmt"
	"os"

	goGithub "github.com/google/go-github/v30/github"
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
