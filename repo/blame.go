package repo

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func (r *Repo) BlameLines(commit *object.Commit) (map[string]int, error) {
	commitIter, err := r.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		log.Fatalf("无法获取提交日志: %v", err)
	}
	defer commitIter.Close()

	blameLines := make(map[string]int)

	// 迭代每个提交对象，累加每个提交者的提交行数
	err = commitIter.ForEach(func(commit *object.Commit) error {
		stats, err := commit.Stats()

		if err != nil {
			return err
		}
		lines := 0
		for _, stat := range stats {
			lines += stat.Addition + stat.Deletion
		}
		blameLines[commit.Author.Email] += lines
		return nil
	})

	fmt.Println(blameLines)

	if err != nil {
		return nil, fmt.Errorf("无法迭代提交: %w", err)
	}

	return blameLines, nil
}
