package main

import (
  "os"
  "errors"
  "io/fs"
  "path"
  "fmt"
  "os/exec"
)

func FileExists(filePath string) bool {
  _, err := os.Stat(filePath)
  return !errors.Is(err, fs.ErrNotExist)
}

func IsFile(filePath string) bool {
  info, _ := os.Stat(filePath)
  return info.Mode().IsRegular()
}

func GzipFileOrDir(filePath string, removeFiles bool) string {
  if IsFile(filePath) {
    return GzipFile(filePath, removeFiles)
  } else {
    return GzipDir(filePath, removeFiles)
  }
}

func GzipFile(filePath string, removeFiles bool) string {
  fileName := path.Base(filePath)
  gzippedPath := fmt.Sprintf("%s.gz", filePath)

  var command *exec.Cmd

  if removeFiles {
    command = exec.Command("gzip", filePath)
  } else {
    // -k to keep the original file.
    command = exec.Command("gzip", "-k", filePath)
  }

  out, err := command.CombinedOutput()

  if err == nil || GzipTest(gzippedPath) {
    return fmt.Sprintf("%s.gz", fileName)
  } else {
    fmt.Println("Gzip failed:")
    fmt.Println(string(out))
    fmt.Println(err.Error())
    // TODO: proper error handling
    return ""
  }
}

func GzipDir(filePath string, removeFiles bool) string {
  fileName := path.Base(filePath)
  fileDir := path.Dir(filePath)
  gzippedPath := fmt.Sprintf("%s.tar.gz", filePath)

  var command *exec.Cmd

  if removeFiles {
    command = exec.Command("tar", "-C", fileDir, "-czf", gzippedPath, fileName, "--remove-files")
  } else {
    /* Use file_path instead of file_name here, so that it creates
     * the archive in the same directory as the file. */
    command = exec.Command("tar", "-C", fileDir, "-czf", gzippedPath, fileName)
  }

  out, err := command.CombinedOutput()

  if err == nil || GzipTest(gzippedPath) {
    return fmt.Sprintf("%s.tar.gz", fileName)
  } else {
    fmt.Println("Gzip failed:")
    fmt.Println(string(out))
    fmt.Println(err.Error())
    // TODO: proper error handling
    return ""
  }
}

func GzipTest(gzippedPath string) bool {
  _, gzipTestErr := exec.Command("gzip", "-t", gzippedPath).Output()
  return gzipTestErr == nil
}

func DestroyFileInArchive(filePath string, archivePath string) bool {
  archiveDir := path.Dir(archivePath)

  command := exec.Command("tar", "-C", archiveDir, "-f", archivePath, "--delete", filePath)

  out, err := command.CombinedOutput()

  if err == nil {
    fmt.Println("Deleted", filePath, "from archive permanently")
    return true
  } else {
    fmt.Println("Deleting from archive failed:")
    fmt.Println(string(out))
    fmt.Println(err.Error())
    // TODO: proper error handling
    return false
  }
}
