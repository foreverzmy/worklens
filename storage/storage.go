package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var home string

type Storage[T any] struct {
	FileName string
}

func init() {
	home = os.Getenv("WORKLENS_HOME")
	if home == "" {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Error fetching user home directory:", err)
			panic(err)
		}
		home = filepath.Join(usr.HomeDir, ".worklens")
	}

	if _, err := os.Stat(home); os.IsNotExist(err) {
		err := os.MkdirAll(home, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			panic(err)
		}
		fmt.Println("Directory created:", home)
	}
}

func (s *Storage[T]) getFilePath() string {
	return filepath.Join(home, s.FileName)
}

func (s *Storage[T]) GetData() (content *T, err error) {
	file := s.getFilePath()

	if _, err := os.Stat(file); err == nil {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read cache file: %w", err)
		}

		content = new(T)
		err = json.Unmarshal(data, content)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
		}

		return content, nil
	}

	return
}

func (s *Storage[T]) SaveData(content T) (err error) {
	file := s.getFilePath()

	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}
	return nil
}

func (s *Storage[T]) DelData() (err error) {
	file := s.getFilePath()

	if _, err := os.Stat(file); err == nil {
		err = os.Remove(file)
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}
