package main

import (
  "fmt"
  "os/exec"
  "strings"
  "path"
)

func ListArchive(archiveName string) {
  if FileExists(archiveName) {
    command := exec.Command("tar", "-t", "-f", archiveName)
    out, err := command.CombinedOutput()

    if err == nil {
      fmt.Println("Files in archive:")
      fmt.Println(string(out))
    } else {
      fmt.Println("An error occurred")
    }
  } else {
    fmt.Println("No archive file in current directory")
  }
}

func ListArchiveTopLevel(archiveName string) {
  if FileExists(archiveName) {
    command := exec.Command("tar", "-t", "-f", archiveName)
    out, err := command.CombinedOutput()

    if err == nil {
      fmt.Println("Top-level files in archive:")
    } else {
      fmt.Println("An error occurred:")
    }

    allFilePaths := strings.Split(string(out), "\n")
    topLevelFiles := filterTopLevelFiles(allFilePaths)
    fmt.Println(strings.Join(topLevelFiles, "\n"))
  } else {
    fmt.Println("No archive file in current directory")
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
