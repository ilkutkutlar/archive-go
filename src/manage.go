package archive

import (
  "path"
  "fmt"
  "errors"
  "os/exec"
)

func AddToArchive(filePath string, archiveName string, removeFiles bool) error {
  fileName := path.Base(filePath)
  fileDir := path.Dir(filePath)

  if !FileExists(filePath) {
    return errors.New(fmt.Sprint("No such file:", filePath))
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
    return nil
  } else {
    errMsg := fmt.Sprint(
      "Adding to archive failed:", "\n",
      string(out), "\n",
      err)
    return errors.New(errMsg)
  }
}

func AddToArchiveGzipped(filePath string, archiveName string, removeFiles bool) (string, error) {
  fileDir := path.Dir(filePath)

  if !FileExists(filePath) {
    return "", errors.New(fmt.Sprint("No such file:", filePath))
  }

  gzippedFileName, err := GzipFileOrDir(filePath, removeFiles)

  if err != nil {
    return "", err
  }

  /* Remove the gzipped file as it is only temporary.
   * There is no option to remove the original file when gzipping. */
   command := exec.Command("tar", "-C", fileDir, "-r", gzippedFileName, "-f", archiveName, "--remove-files")
   out, err := command.CombinedOutput()

   if err == nil {
     return gzippedFileName, nil
   } else {
     errMsg := fmt.Sprint(
       "Adding to archive failed:", "\n",
       string(out), "\n",
       err)
     return "", errors.New(errMsg)
   }
}

func Unarchive(filePath string, archivePath string, removeFiles bool) error {
  archiveDir := path.Dir(archivePath)

  /* Change to archive's directory, so that the unarchived
   * file is placed inside the archive's directory instead of CWD. */
  command := exec.Command("tar", "-C", archiveDir, "-x", "-f", archivePath, filePath)
  out, err := command.CombinedOutput()

  isFileRetrieved := FileExists(fmt.Sprintf("%s/%s", archiveDir, filePath))
  if err != nil || !isFileRetrieved {
    errMsg := fmt.Sprint(
      "Retrieving from archive failed:", "\n",
      string(out), "\n",
      err.Error())
    return errors.New(errMsg)
  }

  if removeFiles {
    return DestroyFileInArchive(filePath, archivePath)
  }

  return nil
}
