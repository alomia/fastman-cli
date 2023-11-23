package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(basePath, fileName string, content ...string) error {
	filePath := filepath.Join(basePath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create the file %s: %w", fileName, err)
	}
	defer file.Close()

	if len(content) > 0 {
		_, err = file.WriteString(content[0])
		if err != nil {
			return fmt.Errorf("failed to write to the file %s: %w", fileName, err)
		}
	}

	return nil
}

func CreateDirectory(basePath, directoryName string) error {
	directoryPath := filepath.Join(basePath, directoryName)

	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return nil
}

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
