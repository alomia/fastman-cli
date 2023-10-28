package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create the file %s: %w", fileName, err)
	}
	defer file.Close()
	return nil
}

func CreateDirectory(basePath, directoryName string) error {
	directoryPath := filepath.Join(basePath, directoryName)
	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func WriteToFile() {}

func FileExists(basePath, fileName string) bool {
	filePath := filepath.Join(basePath, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func DirectoryExists(basePath, directoryName string) bool {
	directoryPath := filepath.Join(basePath, directoryName)
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return false
	}
	return true
}
