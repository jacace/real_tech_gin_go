package main

import "testing"

func TestGetSize(t *testing.T) {
	pr := NewPullRequest()
	size, err := pr.GetSize("")

	if err != nil {
		t.Fatal("Error fecthing PR size")
	}

	if size == 0 {
		t.Fatal("Incorrect PR size")
	}
}
