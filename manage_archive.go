package main

import (
  "path"
  "fmt"
  "os/exec"
)

func AddToArchive(filePath string, archiveName string, removeFiles bool) bool {
  fileName := path.Base(filePath)
  fileDir := path.Dir(filePath)

  if !FileExists(filePath) {
    // TODO: proper error handling
    fmt.Println("No such file:", filePath)
    return false
  }

  var command *exec.Cmd

  if removeFiles {
    /* tar will complain if file path is absolute. To avoid it,
     * change to file's directory and add file by it's basename. */
    command = exec.Command("tar", "-C", fileDir, "-r", fileName, "-f", archiveName, "--remove-files")
  } else {
    command = exec.Command("tar", "-C", fileDir, "-r", fileName, "-f", archiveName)
  }

  out, err := command.CombinedOutput()

  if err == nil {
    fmt.Println(filePath, "added to archive")
    return true
  } else {
    fmt.Println("Adding to archive failed:")
    fmt.Println(string(out))
    fmt.Println(err.Error())
    return false
  }
}

func AddToArchiveGzipped(filePath string, archiveName string, removeFiles bool) bool {
  fileDir := path.Dir(filePath)

  if !FileExists(filePath) {
    // TODO: proper error handling
    fmt.Println("No such file:", filePath)
    return false
  }

  gzippedFileName := GzipFileOrDir(filePath, removeFiles)

  if gzippedFileName == "" {
    /* gzipFileOrDir will print an appropriate error message;
     * just return false here */
    // TODO: proper error handling
    return false
  }

  /* Remove the gzipped file as it is only temporary.
   * There is no option to remove the original file when gzipping. */
   command := exec.Command("tar", "-C", fileDir, "-r", gzippedFileName, "-f", archiveName, "--remove-files")

   out, err := command.CombinedOutput()

   if err == nil {
     fmt.Println(filePath, "added to archive as a gzipped file named", gzippedFileName)
     return true
   } else {
     fmt.Println("Adding to archive failed:")
     fmt.Println(string(out))
     fmt.Println(err.Error())
    // TODO: proper error handling
     return false
   }
}

func Unarchive(filePath string, archivePath string, removeFiles bool) bool {
  archiveDir := path.Dir(archivePath)

  /* Change to archive's directory, so that the unarchived
   * file is placed inside the archive's directory instead of CWD. */
  command := exec.Command("tar", "-C", archiveDir, "-x", "-f", archivePath, filePath)

  out, err := command.CombinedOutput()

  if err == nil && FileExists(fmt.Sprintf("%s/%s", archiveDir, filePath)) {
    fmt.Println("Retrieved", filePath, "from archive")
  } else {
    fmt.Println("Retrieving from archive failed:")
    fmt.Println(string(out))
    fmt.Println(err.Error())
    // TODO: proper error handling
    return false
  }

  if removeFiles {
    return DestroyFileInArchive(filePath, archivePath)
  }

  return true
}
