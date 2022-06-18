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

func (pr *PullRequest) GetSize() (int, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ops := &github.RepositoryListOptions{
		Visibility: "public",
	}
	repos, _, err := client.Repositories.List(ctx, "jacace", ops)

	if err == nil {
		pr.size = len(repos)
		return pr.size, nil
	}

	return 0, err
}
