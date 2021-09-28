package archive

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path"
)

// FileExists returns whether given file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, fs.ErrNotExist)
}

// IsFile returns whether the given file is a regular file
func IsFile(filePath string) bool {
	info, _ := os.Stat(filePath)
	return info.Mode().IsRegular()
}

// DestroyFileInArchive removes the given file from the given archive contents
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

// GzipFileOrDir gzips the given file or directory and returns the gzipped name
func GzipFileOrDir(filePath string, removeFiles bool) (string, error) {
	if IsFile(filePath) {
		return gzipFile(filePath, removeFiles)
	}

	return gzipDir(filePath, removeFiles)
}

func gzipFile(filePath string, removeFiles bool) (string, error) {
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

	if err == nil || gzipTest(gzippedPath) {
		return fileName + ".gz", nil
	}

	errMsg := "Gzip failed:\n" + string(out) + err.Error()
	return "", errors.New(errMsg)
}

func gzipDir(filePath string, removeFiles bool) (string, error) {
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

	if err == nil || gzipTest(gzippedPath) {
		return fileName + ".tar.gz", nil
	}

	errMsg := "Gzip failed:\n" + string(out) + err.Error()
	return "", errors.New(errMsg)
}

func gzipTest(gzippedPath string) bool {
	_, gzipTestErr := exec.Command("gzip", "--test", gzippedPath).Output()
	return gzipTestErr == nil
}
