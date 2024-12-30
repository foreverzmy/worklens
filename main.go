package main

import (
	"fmt"
	"log"

	"github.com/foreverzmy/worklens/repo"
)

const repoRoot = ""

func main() {

	r, err := repo.PlainOpen(repoRoot)
	if err != nil {
		panic(err)
	}

	commit, _, err := r.GetBranchFiles("main")
	if err != nil {
		panic(err)
	}

	blameLines, err := r.BlameLines(commit)
	if err != nil {
		log.Fatalf("Failed to get pure blame: %v", err)
	}
	fmt.Println(blameLines)
}
