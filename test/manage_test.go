package archive

import (
  "testing"
  "os"
  "path"
  "strings"
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

  expected := strings.Join([]string{
    path.Base(testFile1),
    path.Base(testDir1) + "/",
    path.Base(testDir1) + "/" + path.Base(testFile2),
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileExists(t, testFile1)
  assertFileExists(t, testDir1)
}

func TestAddToArchiveFilesWithSpaces(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := tempDir + "/test file 1.txt"
  testDir1 := tempDir + "/test dir"
  testFile2 := testDir1 + "/test file 2.txt"

  os.Create(testFile1)
  os.Mkdir(testDir1, 0755)
  os.Create(testFile2)

  archive.AddToArchive(testFile1, TEST_ARCHIVE, false)
  archive.AddToArchive(testDir1, TEST_ARCHIVE, false)

  expected := strings.Join([]string{
    path.Base(testFile1),
    path.Base(testDir1) + "/",
    path.Base(testDir1) + "/" + path.Base(testFile2),
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileExists(t, testFile1)
  assertFileExists(t, testDir1)
}

func TestAddToArchiveAndRemove(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := createTestFile(tempDir, 1)
  testDir1 := createTestDir(tempDir, 1)
  testFile2 := createTestFile(testDir1, 2)

  archive.AddToArchive(testFile1, TEST_ARCHIVE, true)
  archive.AddToArchive(testDir1, TEST_ARCHIVE, true)

  expected := strings.Join([]string{
    path.Base(testFile1),
    path.Base(testDir1) + "/",
    path.Base(testDir1) + "/" + path.Base(testFile2),
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileDoesNotExist(t, testFile1)
  assertFileDoesNotExist(t, testDir1)
}

func TestAddToArchiveFilesWithSpacesAndRemove(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := tempDir + "/test file 1.txt"
  testDir1 := tempDir + "/test dir"
  testFile2 := testDir1 + "/test file 2.txt"

  os.Create(testFile1)
  os.Mkdir(testDir1, 0755)
  os.Create(testFile2)

  archive.AddToArchive(testFile1, TEST_ARCHIVE, true)
  archive.AddToArchive(testDir1, TEST_ARCHIVE, true)

  expected := strings.Join([]string{
    path.Base(testFile1),
    path.Base(testDir1) + "/",
    path.Base(testDir1) + "/" + path.Base(testFile2),
  }, "\n")

  assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
  assertFileDoesNotExist(t, testFile1)
  assertFileDoesNotExist(t, testDir1)
}

func TestErrorHandledCorrectlyDuringArchiving(t *testing.T) {
  t.SkipNow()
}

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

func TestUnarchiveFilesFromArchive(t *testing.T) {
  t.SkipNow()
}

func TestUnarchiveFilesWithSpacesInNameFromArchive(t *testing.T) {
  t.SkipNow()
}

func TestUnarchiveFilesFromArchiveAndRemove(t *testing.T) {
  t.SkipNow()
}

func TestUnarchiveFilesWithSpacesInNameFromArchiveAndRemove(t *testing.T) {
  t.SkipNow()
}

func cleanup() {
  os.Remove(TEST_ARCHIVE)
}
