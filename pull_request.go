package main

import (
	"context"

	"github.com/google/go-github/v45/github"

	"golang.org/x/oauth2"
)

type PullRequest struct {
	user string
	size int
	id   string
}

func NewPullRequest() *PullRequest {
	return &PullRequest{
		user: "",
		id:   "",
		size: 0,
	}
}

func (pr *PullRequest) GetSize(personalAccessToken string) (int, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalAccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ops := &github.PullRequestListOptions{
		State:     "Closed",
		Base:      "main",
		Sort:      "created",
		Direction: "asc",
	}

	prs, _, err := client.PullRequests.List(ctx, "jacace", "real_tech_gin_go", ops)

	if err == nil {
		pr.size = len(prs)
		return pr.size, nil
	}

	return 0, err
}
