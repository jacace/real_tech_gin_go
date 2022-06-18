package main

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

func (prs *PullRequest) getSize() error {
	prs.size = 0 // place holder to call github API
	return nil
}
