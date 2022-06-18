package main

import "fmt"

func main() {
	fmt.Println("Initial commit")
	pr := NewPullRequest()
	pr.getSize()
}
