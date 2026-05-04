package main

import (
	"fmt"
	"log"

	"git-whereami/git"
)

func main() {
	branch, err := git.CurrentBranch()
	if err != nil || branch == "" {
		log.Fatal("not on a branch (detached HEAD?)")
	}

	base, err := git.DefaultBranch()
	if err != nil {
		// fallback to local main or master
		if _, err2 := git.MergeBase("main"); err2 == nil {
			base = "main"
		} else if _, err2 := git.MergeBase("master"); err2 == nil {
			base = "master"
		} else {
			log.Fatal("could not determine base branch (no origin, main, or master)")
		}
	}


	ahead, behind, err := git.AheadBehind(base)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Printf("Base branch   : %s\n", base)
	fmt.Printf("Current branch: %s\n", branch)
	fmt.Printf("Ahead         : %s\n", ahead)
	fmt.Printf("Behind        : %s\n", behind)
	graph, err := git.Graph(base)
if err == nil && graph != "" {
	fmt.Println("\nYou are here:")
	fmt.Println(graph)
}
}
