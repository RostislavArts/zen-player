package fsutil

import (
	"os"
	"fmt"
	"path/filepath"
)

func isPathDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func ParseFilesFromDir(path string) ([]string, error) {
	var fileList []string

	isDir, err := isPathDir(path)
	if err != nil {
		return nil, err
	}

	if isDir {
		directory, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, file := range directory {
			filename := file.Name()

			// Check for a file to be a proper format
			if filepath.Ext(filename) == ".mp3" ||
			filepath.Ext(filename) == ".wav" ||
			filepath.Ext(filename) == ".flac" {
				fileList = append(fileList, fmt.Sprintf("%s/%s", path, filename))
			}
		}

		return fileList, nil
	} else {
		fileList = append(fileList, path)
		return fileList, nil
	}
}

