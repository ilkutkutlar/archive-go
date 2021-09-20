package archive

import (
  "os"
  "fmt"
  "os/exec"
  "errors"
  "io/fs"
  "strings"
  "testing"
)

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
