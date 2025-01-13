package main

import (
	"log"
	"os"
	"time"

	"github.com/foreverzmy/worklens/repo"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const repoRoot = "/Users/bytedance/Workspace/ditom"
const dateLimit = "2025-01-01"

func main() {
	// 创建或打开一个日志文件
	os.Remove("logs/app.log")
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer file.Close()
	log.SetFlags(0)
	// 设置日志输出到文件
	log.SetOutput(file)

	r, err := repo.PlainOpen(repoRoot)
	if err != nil {
		log.Println("Error opening repository:", err)
		return
	}

	head, err := r.Head()

	if err != nil {
		log.Println("Error get Head Branch:", err)
	}
	log.Printf("head: %s\n", head.String())

	commit, _, err := r.GetBranchFiles(head.Name().Short())
	if err != nil {
		log.Println("Error get Head branch files:", err)
		return
	}

	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		log.Println("Error getting commit log:", err)
		return
	}

	limitTime, _ := time.Parse("2006-01-02", dateLimit)

	log.Println("Commits before", dateLimit, ":")

	count := 0
	err = commitIter.ForEach(func(c *object.Commit) error {
		if c.Committer.When.Before(limitTime) {
			count++
			log.Printf("%d:\nCommit: %s\nAuthor: %s\nDate: %s\nMessage: %s\n\n", count, c.Hash, c.Author.Email, c.Committer.When, c.Message)
		}
		return nil
	})

	if err != nil {
		log.Println("Error iterating through commits:", err)
	}
}
