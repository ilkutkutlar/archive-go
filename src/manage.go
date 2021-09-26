package archive

import (
	"errors"
	"path"
)

func AddToArchive(filePath string, archiveName string, removeFiles bool) error {
	if !FileExists(filePath) {
		return errors.New("No such file: " + filePath)
	}

	out, err := tarAddToArchive(filePath, archiveName, removeFiles)

	if err != nil {
		errMsg := "Adding to archive failed:\n" + out + err.Error()
		return errors.New(errMsg)
	}

	return nil
}

func AddToArchiveGzipped(filePath string, archiveName string, removeFiles bool) (string, error) {
	fileDir := path.Dir(filePath)

	if !FileExists(filePath) {
		return "", errors.New("No such file: " + filePath)
	}

	gzippedFileName, err := GzipFileOrDir(filePath, removeFiles)
	gzippedFilePath := path.Join(fileDir, gzippedFileName)

	if err != nil {
		return "", err
	}

	/* Remove the gzipped file as it is only temporary.
	 * There is no option in gzip to remove the original file when gzipping. */
	out, err := tarAddToArchive(gzippedFilePath, archiveName, true)

	if err != nil {
		errMsg := "Adding to archive failed:\n" + out + err.Error()
		return "", errors.New(errMsg)
	}

	return gzippedFileName, nil
}

func Unarchive(filePath string, archiveName string, removeFiles bool) error {
	archiveDir := path.Dir(archiveName)

	out, err := tarUnarchive(filePath, archiveName)

	isFileRetrieved := FileExists(path.Join(archiveDir, filePath))
	if err != nil || !isFileRetrieved {
		errMsg := "Retrieving from archive failed:\n" + out + err.Error()
		return errors.New(errMsg)
	}

	if removeFiles {
		return DestroyFileInArchive(filePath, archiveName)
	}

	return nil
}
