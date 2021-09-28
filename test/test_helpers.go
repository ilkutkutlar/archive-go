package archive

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

const (
	testArchive    = "test_archive.tar"
	dummyArchive1 = "fixtures/dummy_archive.tar"
	dummyArchive2 = "fixtures/dummy_archive2.tar"
)

var (
	testFile1 string
	testFile2 string
	testFile3 string
	testFile4 string
	testDir1  string
	testDir2  string
)

func cleanup() {
	os.Remove(testArchive)
}

func createTestFiles(t *testing.T) {
	// temp dir is removed after the unit test function
	// returns, so no need to remove it in cleanup
	tempDir := t.TempDir()

	testDir1 = path.Join(tempDir, "test_dir1")
	testFile1 = path.Join(tempDir, "test1.txt")
	testFile2 = path.Join(testDir1, "test2.txt")

	testDir2 = path.Join(tempDir, "test dir")
	testFile3 = path.Join(tempDir, "test file 1.txt")
	testFile4 = path.Join(testDir2, "test file 2.txt")

	testFiles := []string{testFile1, testFile2, testFile3, testFile4}
	testDirs := []string{testDir1, testDir2}

	for _, testDir := range testDirs {
		os.Mkdir(testDir, 0755)
	}

	for _, testFile := range testFiles {
		os.Create(testFile)
	}
}

func assertStringEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected: '%s'\nActual: '%s'", expected, actual)
	}
}

func assertFileExists(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)

	if errors.Is(err, fs.ErrNotExist) {
		t.Fail()
	}
}

func assertFileDoesNotExist(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)

	if err == nil || !errors.Is(err, fs.ErrNotExist) {
		t.Fail()
	}
}

func assertArchiveContentsEqual(t *testing.T, archivePath string, expectedArchiveFiles string) {
	actual := getArchiveFiles(archivePath)

	if expectedArchiveFiles != actual {
		t.Errorf("Expected: '%s'\nActual: '%s'", expectedArchiveFiles, actual)
	}
}

func assertNil(t *testing.T, object interface{}) {
	if object != nil {
		t.Fail()
	}
}

func getArchiveFiles(archivePath string) string {
	command := exec.Command("tar", "--list", "--file", archivePath)
	out, _ := command.CombinedOutput()
	return strings.TrimSpace(string(out))
}

func copyFile(srcPath string, destPath string) error {
	in, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	return out.Close()
}
