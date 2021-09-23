package archive

import (
  "fmt"
  "os/exec"
  "strings"
  "path"
  "errors"
)

func ListArchive(archiveName string) (string, error) {
  if !FileExists(archiveName) {
    return "", errors.New("No archive file in current directory")
  }

  command := exec.Command("tar", "-t", "-f", archiveName)
  out, err := command.CombinedOutput()

  if err == nil {
    return fmt.Sprint(string(out)), nil
  } else {
    return "", errors.New("An error occurred")
  }
}

func ListArchiveTopLevel(archiveName string) (string, error) {
  if !FileExists(archiveName) {
    return "", errors.New("No archive file in current directory")
  }

  command := exec.Command("tar", "-t", "-f", archiveName)
  out, err := command.CombinedOutput()

  if err == nil {
    allFilePaths := strings.Split(string(out), "\n")
    topLevelFiles := filterTopLevelFiles(allFilePaths)
    return strings.Join(topLevelFiles, "\n"), nil
  } else {
    errMsg := fmt.Sprint("An error occurred:", string(out))
    return "", errors.New(errMsg)
  }
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
  return path.Dir(filePath) == path.Base(filePath)
}
