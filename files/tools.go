package files

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteFilesToText(details []FileInfo, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for index, detail := range details {
		_, err := fmt.Fprintf(file, "Index: %d, Type: %s, Path: %s, Name: %s\n", index, detail.FileType, detail.FilePath, detail.FileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteFilesToCSV(details []FileInfo, fileCSVPath string) error {
	file, err := os.Create(fileCSVPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Index", "FileType", "FilePath", "FileName", "FileSize"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for index, detail := range details {
		record := []string{
			fmt.Sprintf("%d", index),
			detail.FileType,
			detail.FilePath,
			detail.FileName,
			fmt.Sprintf("%d", detail.FileSize),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
