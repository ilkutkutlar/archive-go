package archive

import (
	"errors"
	"path"
	"strings"
)

// ListArchive returns a new line separated string of all archive contents recursively.
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

// ListArchiveTopLevel returns a new line separated string of archive contents without a parent directory.
func ListArchiveTopLevel(archiveName string) (string, error) {
	if !FileExists(archiveName) {
		return "", errors.New("No archive file in current directory")
	}

	out, err := tarGetContents(archiveName)

	if err != nil {
		errMsg := "An error occurred: " + out
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
