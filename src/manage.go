package archive

import (
	"errors"
	"fmt"
	"path"
)

func AddToArchive(filePath string, archiveName string, removeFiles bool) error {
  if !FileExists(filePath) {
    return errors.New(fmt.Sprint("No such file:", filePath))
  }

  out, err := tarAddToArchive(filePath, archiveName, removeFiles)

  if err == nil {
    return nil
  } else {
    errMsg := fmt.Sprint(
      "Adding to archive failed:", "\n",
      out, "\n",
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
  gzippedFilePath := path.Join(fileDir, gzippedFileName)

  if err != nil {
    return "", err
  }

  /* Remove the gzipped file as it is only temporary.
   * There is no option in gzip to remove the original file when gzipping. */
   out, err := tarAddToArchive(gzippedFilePath, archiveName, true)

   if err == nil {
     return gzippedFileName, nil
   } else {
     errMsg := fmt.Sprint(
       "Adding to archive failed:", "\n",
       out, "\n",
       err)
     return "", errors.New(errMsg)
   }
}

func Unarchive(filePath string, archiveName string, removeFiles bool) error {
  archiveDir := path.Dir(archiveName)

  out, err := tarUnarchive(filePath, archiveName)

  isFileRetrieved := FileExists(fmt.Sprintf("%s/%s", archiveDir, filePath))
  if err != nil || !isFileRetrieved {
    errMsg := fmt.Sprint(
      "Retrieving from archive failed:", "\n",
      out, "\n",
      err.Error())
    return errors.New(errMsg)
  }

  if removeFiles {
    return DestroyFileInArchive(filePath, archiveName)
  }

  return nil
}
