package files

import (
	"bufio"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	FilePath string
	FileName string
	FileType string
	FileSize int64
	FileMode os.FileMode
	ModTime  time.Time
	IsDir    bool
}

func readGitignore(root string) (*GitIgnore, error) {
	gitignorePath := filepath.Join(root, ".gitignore")
	file, err := os.Open(gitignorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return CompileIgnoreLines(), nil
		}
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	lines = append(lines, ".git/")

	return CompileIgnoreLines(lines...), nil
}

// ListFiles recursively lists all files in the directory, respecting .gitignore rules
func ListFiles(root string) ([]FileInfo, error) {
	ignore, err := readGitignore(root)
	if err != nil {
		return nil, err
	}

	var files []FileInfo

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == root {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if ignore.MatchesPath(relPath) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
			files = append(files, FileInfo{
				FilePath: path,
				FileName: info.Name(),
				FileType: filepath.Ext(info.Name()),
				FileSize: info.Size(),
				FileMode: info.Mode(),
				ModTime:  info.ModTime(),
				IsDir:    info.IsDir(),
			})
		}

		return nil
	})

	return files, err
}
