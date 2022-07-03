package main

import "testing"

func TestGetSize(t *testing.T) {
	client := NewGitHubAPI("")
	pr, err := client.GetSize("jacace", "real_tech_gin_go")

	if err != nil {
		t.Fatal("Error fecthing PR size")
	}

	if pr.size == 0 {
		t.Fatal("Incorrect PR size")
	}
}
