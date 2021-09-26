package archive

import (
	"os/exec"
	"path"
)

func tarAddToArchive(filePath string, archiveName string, removeFiles bool) (string, error) {
	fileName := path.Base(filePath)
	fileDir := path.Dir(filePath)

	/* tar will complain if file path is absolute. To avoid it,
	 * change to file's directory and add file by it's basename. */
	args := []string{
		"--directory", fileDir,
		"--file", archiveName,
		"--append", fileName}

	if removeFiles {
		args = append(args, "--remove-files")
	}

	command := exec.Command("tar", args...)
	out, err := command.CombinedOutput()
	return string(out), err
}

func tarUnarchive(filePath string, archiveName string) (string, error) {
	archiveDir := path.Dir(archiveName)

	/* Change to archive's directory, so that the unarchived
	 * file is placed inside the archive's directory instead of CWD. */
	args := []string{
		"--directory", archiveDir,
		"--file", archiveName,
		"--extract",
		filePath}
	command := exec.Command("tar", args...)
	out, err := command.CombinedOutput()
	return string(out), err
}

func tarGetContents(archiveName string) (string, error) {
	command := exec.Command("tar", "--file", archiveName, "--list")
	out, err := command.CombinedOutput()
	return string(out), err
}
