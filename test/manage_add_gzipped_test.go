package archive

import (
  "testing"
  "path"
  "strings"
  archive "example.com/archive/src"
)

func TestAddToArchiveGzipped(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := createTestFile(tempDir, 1)
  testDir1 := createTestDir(tempDir, 1)
  createTestFile(testDir1, 2)

  archive.AddToArchiveGzipped(testFile1, TEST_ARCHIVE, false)
  archive.AddToArchiveGzipped(testDir1, TEST_ARCHIVE, false)

  expected := strings.Join([]string{
    path.Base(testFile1) + ".gz",
    path.Base(testDir1) + ".tar.gz",
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileExists(t, testFile1)
  assertFileExists(t, testDir1)
  // The gzipped files are temporary: must be removed after adding.
  assertFileDoesNotExist(t, testFile1 + ".gz")
  assertFileDoesNotExist(t, testDir1 + ".tar.gz")
}

func TestAddToArchiveGzippedAndRemove(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := createTestFile(tempDir, 1)
  testDir1 := createTestDir(tempDir, 1)
  testFile2 := createTestFile(testDir1, 2)

  archive.AddToArchiveGzipped(testFile1, TEST_ARCHIVE, true)
  archive.AddToArchiveGzipped(testDir1, TEST_ARCHIVE, true)

  expected := strings.Join([]string{
    path.Base(testFile1) + ".gz",
    path.Base(testDir1) + ".tar.gz",
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileDoesNotExist(t, testFile1)
  assertFileDoesNotExist(t, testDir1)
  assertFileDoesNotExist(t, testFile2)
  // The gzipped files are temporary: must be removed after adding.
  assertFileDoesNotExist(t, testFile1 + ".gz")
  assertFileDoesNotExist(t, testDir1 + ".tar.gz")
}

func TestErrorHandledCorrectlyDuringArchivingGzipped(t *testing.T) {
  t.SkipNow()
}
