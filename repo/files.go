package repo

import (
	"log"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type FileInfo struct {
	*object.File
	FilePath string
	FileName string
	FileType string
}

func (r *Repo) GetBranchFiles(branchName string) (*object.Commit, []*FileInfo, error) {
	var files []*FileInfo
	// 获取参考（reference）以切换到目标分支
	refName := plumbing.NewBranchReferenceName(branchName)
	ref, err := r.Reference(refName, true)
	if err != nil {
		log.Fatalf("Failed to get reference: %s", err)
		return nil, nil, err
	}

	// 获取目标分支的commit对象
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatalf("Failed to get commit object: %s", err)
		return nil, nil, err
	}

	// 获取目标分支的树(tree)对象
	tree, err := commit.Tree()
	if err != nil {
		log.Fatalf("Failed to get tree object: %s", err)
		return nil, nil, err
	}
	// 遍历树对象以列出所有文件
	err = tree.Files().ForEach(func(f *object.File) error {
		files = append(files, &FileInfo{
			File:     f,
			FilePath: filepath.Dir(f.Name),
			FileName: filepath.Base(f.Name),
			FileType: filepath.Ext(f.Name),
		})
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to list files: %s", err)
		return nil, nil, err
	}

	return commit, files, nil
}
