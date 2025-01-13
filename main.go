package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/foreverzmy/worklens/repo"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"

	ignore "github.com/sabhiram/go-gitignore"
)

const dateLimit = "2025-01-01"

type BlameState struct {
	Addition int
	Deletion int
}

func main() {
	repoFlag := flag.String("repo", "", "")
	flag.Parse()

	ig, _ := ignore.CompileIgnoreFile("fileignore")

	repoRoot := *repoFlag
	if repoRoot == "" {
		log.Fatalf("仓库目录未设置")
	}

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
	fileName := fmt.Sprintf("cache/%s.yaml", strings.ReplaceAll(head.Name().Short(), "/", "_"))
	os.Remove(fileName)
	headFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer headFile.Close()

	writer := bufio.NewWriter(headFile)
	_, err = writer.WriteString(fmt.Sprintf("branch: %s\n", head.Name().Short()))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

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

	blameLines := make(map[string]*BlameState)

	limitTime, _ := time.Parse("2006-01-02", dateLimit)

	log.Println("Commits before", dateLimit, ":")

	writer.WriteString("commits:\n")
	count := 0
	err = commitIter.ForEach(func(c *object.Commit) error {
		if c.Committer.When.Before(limitTime) {
			count++
			writer.WriteString(fmt.Sprintf("  - No: %d\n", count))
			writer.WriteString(fmt.Sprintf("    hash: \"%s\"\n", c.Hash))
			writer.WriteString(fmt.Sprintf("    author: \"%s\"\n", c.Author.Email))
			writer.WriteString(fmt.Sprintf("    date: \"%s\"\n", c.Committer.When.String()))
			commitMessage := strings.TrimSpace(c.Message)
			multiLineString := strings.ReplaceAll(commitMessage, "\r\n", "\n")
			singleLineMessage := strings.ReplaceAll(multiLineString, "\n", "\\n")
			writer.WriteString(fmt.Sprintf("    message: \"%s\"\n", singleLineMessage))

			if blameLines[c.Author.Email] == nil {
				blameLines[c.Author.Email] = &BlameState{0, 0}
			}

			stats, err := c.Stats()
			if err != nil {
				return err
			}
			addition := 0
			deletion := 0
			writer.WriteString("    stat:\n")
			for _, stat := range stats {
				names := strings.Split(stat.Name, " => ")
				if !ig.MatchesPath(names[0]) {
					writer.WriteString(fmt.Sprintf("      - \"%s\"\n", stat.Name))
					addition += stat.Addition
					deletion += stat.Deletion
				}
			}
			blameLines[c.Author.Email].Addition += addition
			blameLines[c.Author.Email].Deletion += deletion
			writer.WriteString(fmt.Sprintf("    addition: %d\n", addition))
			writer.WriteString(fmt.Sprintf("    deletion: %d\n", deletion))
		}
		return nil
	})

	if err != nil {
		log.Println("Error iterating through commits:", err)
	}

	writer.WriteString("blame:\n")
	for email, state := range blameLines {
		writer.WriteString(fmt.Sprintf("  - email: \"%s\"\n", email))
		writer.WriteString(fmt.Sprintf("    addition: %d\n", state.Addition))
		writer.WriteString(fmt.Sprintf("    deletion: %d\n", state.Deletion))
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
	}
}
