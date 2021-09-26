package archive

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

func ListArchive(archiveName string) (string, error) {
	if !FileExists(archiveName) {
		return "", errors.New("No archive file in current directory")
	}

	out, err := tarGetContents(archiveName)

	if err != nil {
		return "", errors.New("An error occurred")
	}

	return out, nil
}

func ListArchiveTopLevel(archiveName string) (string, error) {
	if !FileExists(archiveName) {
		return "", errors.New("No archive file in current directory")
	}

	out, err := tarGetContents(archiveName)

	if err != nil {
		errMsg := fmt.Sprint("An error occurred:", out)
		return "", errors.New(errMsg)
	}

	allFilePaths := strings.Split(out, "\n")
	topLevelFiles := filterTopLevelFiles(allFilePaths)
	return strings.Join(topLevelFiles, "\n"), nil
}

func filterTopLevelFiles(allFilePaths []string) []string {
	var topLevelFiles []string

	for _, filePath := range allFilePaths {
		if isTopLevel(filePath) {
			topLevelFiles = append(topLevelFiles, filePath)
		}
	}

	return topLevelFiles
}

func isTopLevel(filePath string) bool {
	isTopLevelDir := path.Dir(filePath) == path.Base(filePath)
	isTopLevelFile := path.Dir(filePath) == "."
	return isTopLevelDir || isTopLevelFile
}
