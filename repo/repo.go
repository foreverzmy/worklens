package repo

import (
	"github.com/go-git/go-git/v5"
)

type Repo struct {
	*git.Repository
	Path string
}

func PlainOpen(repoPath string) (*Repo, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	repo := &Repo{
		Repository: r,
		Path:       repoPath,
	}

	return repo, nil
}
