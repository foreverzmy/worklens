package repo

import (
	"fmt"
	"log"

	"github.com/foreverzmy/worklens/storage"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func (r *Repo) PureBlameLines(commit *object.Commit, files []*FileInfo) (map[string]map[string]int, error) {

	cache := storage.Storage[map[string]map[string]int]{FileName: fmt.Sprintf("blame.%s.json", commit.Hash.String())}

	blameLines := make(map[string]map[string]int)

	data, err := cache.GetData()
	if err == nil && *data != nil {
		blameLines = *data
	}

	for _, file := range files {
		if _, ok := blameLines[file.Name]; ok {
			continue
		}
		blameResult, err := git.Blame(commit, file.Name)
		if err != nil {
			log.Fatalf("Failed to get blame for file: %v", err)
		}
		// 存储每个作者的提交行数
		fileAuthorLines := make(map[string]int)
		for _, line := range blameResult.Lines {
			fileAuthorLines[line.Author] += 1
		}

		blameLines[file.Name] = fileAuthorLines

		err = cache.SaveData(blameLines)
		if err != nil {
			log.Fatalf("Failed to save blame to file: %v", err)
		}
	}
	return blameLines, nil
}
