package archive

import (
  "testing"
  "os"
  "path"
  "strings"
  "gotest.tools/v3/assert"
  archive "example.com/archive/src"
)

const TEST_ARCHIVE = "test_archive.tar"

func TestAddToArchive(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()

  testFile1 := createTestFile(tempDir, 1)
  testDir1 := createTestDir(tempDir, 1)
  testFile2 := createTestFile(testDir1, 2)

  archive.AddToArchive(testFile1, TEST_ARCHIVE, false)
  archive.AddToArchive(testDir1, TEST_ARCHIVE, false)

  actual := getArchiveFiles(TEST_ARCHIVE)
  expected := strings.Join([]string{
    path.Base(testFile1),
    path.Base(testDir1) + "/",
    path.Base(testDir1) + "/" + path.Base(testFile2),
  }, "\n")

  assert.Equal(t, expected, actual)
  assertFileExists(t, testFile1)
  assertFileExists(t, testDir1)
}

func cleanup() {
  os.Remove(TEST_ARCHIVE)
}
