package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v45/github"

	"golang.org/x/oauth2"
)

type PullRequest struct {
	leadTimeHrs float64 //date and time for each pull request when opened, and then, when it’s merged
	timeToMerge float64 //timestamp of the oldest commit of a branch minus the timestamp of the merge commit
	size        int     //average of total lines of code added plus the total lines of code removed
	discussions int     //number of comments and reactions per pull request
}

type GitHubAPI struct {
	token string
}

func NewGitHubAPI(token string) *GitHubAPI {
	return &GitHubAPI{
		token: token,
	}
}

func (pr *GitHubAPI) GetSize(user string, repo string) (*PullRequest, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pr.token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ops := &github.PullRequestListOptions{
		State:     "open", //closed
		Base:      "main",
		Sort:      "created",
		Direction: "asc",
	}

	prs, _, err := client.PullRequests.List(ctx, user, repo, ops)

	if err == nil {
		for _, a := range prs {
			//a.GetClosedAt()
			urls := a.GetCommitsURL()
			fmt.Println("Commits URL:" + urls)

			pr := &PullRequest{
				leadTimeHrs: (a.GetMergedAt().Sub(a.GetCreatedAt())).Hours(),
			}

			return pr, nil
		}
	}

	return nil, err
}
