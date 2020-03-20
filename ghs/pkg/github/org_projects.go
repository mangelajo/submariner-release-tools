package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v30/github"
)

func (gh *GitHub) GetProjects(owner string) ([]*github.Project, *github.Response, error) {

	opt := &github.ProjectListOptions{State: "open", ListOptions: github.ListOptions{PerPage: 100}}

	return gh.client.Organizations.ListProjects(context.Background(), owner, opt)
}

func (gh *GitHub) GetProjectIDs(owner string, urls []string) (ids []int64, err error) {
	projects, _, err := gh.GetProjects(owner)
	if err != nil {
		return
	}

	for _, url := range urls {
		id := findProjectWithUrl(projects, url)
		if id == nil {
			return nil, fmt.Errorf("Could not find project with url %s in owner %s", url, owner)
		}
		ids = append(ids, *id)
	}
	return
}

func findProjectWithUrl(projects []*github.Project, url string) *int64 {
	for _, project := range projects {
		if *project.HTMLURL == url {
			return project.ID
		}
	}
	return nil
}

func (gh *GitHub) GetProjectColumns(projectID int64) ([]*github.ProjectColumn, *github.Response, error) {

	opt := &github.ListOptions{PerPage: 100}
	return gh.client.Projects.ListProjectColumns(context.Background(), projectID, opt)

}

func (gh *GitHub) GetProjectColumn(projectID int64, columnName string) (int64, error) {

	cols, _, err := gh.GetProjectColumns(projectID)
	if err != nil {
		return -1, err
	}
	for _, col := range cols {
		if strings.ToLower(*col.Name) == strings.ToLower(columnName) {
			return *col.ID, nil
		}
	}
	return -1, fmt.Errorf("Column %s not found for project id %d", columnName, projectID)
}

func (gh *GitHub) GetColumnCards(columnID int64) ([]*github.ProjectCard, *github.Response, error) {
	opt := &github.ProjectCardListOptions{ListOptions: github.ListOptions{PerPage: 100}}
	return gh.client.Projects.ListProjectCards(context.Background(), columnID, opt)
}

func (gh *GitHub) MoveCardToAnotherProjectColumn(card *github.ProjectCard, columnID int64) (*github.ProjectCard, *github.Response, error) {

	cardOptions, err := ProjectCardOption(card)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(cardOptions)
	newCard, resp, err := gh.client.Projects.CreateProjectCard(context.Background(), columnID, &cardOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("creating new card: err=%s resp=%+v columnID=%d, cardOptions=%+v", err, resp,
			columnID, cardOptions)
	}

	_, err = gh.client.Projects.DeleteProjectCard(context.Background(), *card.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("deleting old card: %s", err)
	}

	return newCard, resp, nil

}

func ProjectCardOption(card *github.ProjectCard) (cardOpts github.ProjectCardOptions, err error) {

	if card.Note != nil {
		cardOpts.Note = *card.Note
		return
	}
	//           0     1 2                3         4          5         6    7
	// example: https://api.github.com/repos/submariner-io/lighthouse/issues/81
	path := strings.Split(*card.ContentURL, "/")

	contentId, err := strconv.ParseInt(path[7], 10, 64)
	if err != nil {
		return cardOpts, fmt.Errorf("Unable to parse content ID from URL: %s", err)
	}

	contentType := path[6]

	if contentType == "issues" {
		cardOpts.ContentType = "Issue"
	} else if contentType == "pull" {
		cardOpts.ContentType = "PullRequest"
	}
	cardOpts.ContentID = contentId
	return
}
