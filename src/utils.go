package archive

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, fs.ErrNotExist)
}

func IsFile(filePath string) bool {
	info, _ := os.Stat(filePath)
	return info.Mode().IsRegular()
}

func DestroyFileInArchive(filePath string, archivePath string) error {
	archiveDir := path.Dir(archivePath)

	command := exec.Command("tar",
		"--directory", archiveDir,
		"--file", archivePath,
		"--delete", filePath)

	out, err := command.CombinedOutput()

	if err != nil {
		errMsg := "Deleting from archive failed:\n" + string(out) + err.Error()
		return errors.New(errMsg)
	}

	return nil
}

func GzipFileOrDir(filePath string, removeFiles bool) (string, error) {
	if IsFile(filePath) {
		return GzipFile(filePath, removeFiles)
	} else {
		return GzipDir(filePath, removeFiles)
	}
}

func GzipFile(filePath string, removeFiles bool) (string, error) {
	fileName := path.Base(filePath)
	gzippedPath := filePath + ".gz"

	var command *exec.Cmd

	if removeFiles {
		command = exec.Command("gzip", filePath)
	} else {
		// keep the original file.
		command = exec.Command("gzip", "--keep", filePath)
	}

	out, err := command.CombinedOutput()

	if err == nil || GzipTest(gzippedPath) {
		return fileName + ".gz", nil
	} else {
		errMsg := "Gzip failed:\n" + string(out) + err.Error()
		return "", errors.New(errMsg)
	}
}

func GzipDir(filePath string, removeFiles bool) (string, error) {
	fileName := path.Base(filePath)
	fileDir := path.Dir(filePath)
	gzippedPath := filePath + ".tar.gz"

	/* Use file_path instead of fileName here, so that it creates
	 * the archive in the same directory as the file. */
	args := []string{
		"--directory", fileDir,
		"--file", gzippedPath,
		"--create",
		"--gzip",
		fileName}

	if removeFiles {
		args = append(args, "--remove-files")
	}

	command := exec.Command("tar", args...)
	out, err := command.CombinedOutput()

	if err == nil || GzipTest(gzippedPath) {
		return fileName + ".tar.gz", nil
	} else {
		errMsg := "Gzip failed:\n" + string(out) + err.Error()
		return "", errors.New(errMsg)
	}
}

func GzipTest(gzippedPath string) bool {
	_, gzipTestErr := exec.Command("gzip", "--test", gzippedPath).Output()
	return gzipTestErr == nil
}
