package main

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/v45/github"

	"net/http"

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
			resGet, err := http.Get(a.GetCommitsURL())

			if err == nil {
				respBody, err := ioutil.ReadAll(resGet.Body)
				if err == nil {
					str := string(respBody)
					if str != "" {
						str = ""
					}
				}
				resGet.Body.Close()
			}

			pr := &PullRequest{
				leadTimeHrs: (a.GetMergedAt().Sub(a.GetCreatedAt())).Hours(),
				timeToMerge: 0,
				discussions: a.GetComments(),
				size:        0,
			}

			return pr, nil
		}
	}

	return nil, err
}
