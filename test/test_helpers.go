package archive

import (
  "os"
  "fmt"
  "os/exec"
  "errors"
  "io/fs"
  "strings"
  "testing"
  "io"
)

const TEST_ARCHIVE = "test_archive.tar"

func cleanup() {
  os.Remove(TEST_ARCHIVE)
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

  if actual != expectedArchiveFiles {
    t.Fail()
  }
}

func getArchiveFiles(archivePath string) string {
    command := exec.Command("tar", "-t", "-f", archivePath)
    out, _ := command.CombinedOutput()
    return strings.TrimSpace(string(out))
}

func createTestFile(dir string, number int) string {
  testFile := fmt.Sprintf("%s/test%d.txt", dir , number)
  os.Create(testFile)
  return testFile
}

func createTestDir(dir string, number int) string {
  testDir:= fmt.Sprintf("%s/test_dir%d.txt", dir , number)
  os.Mkdir(testDir, 0755)
  return testDir
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
