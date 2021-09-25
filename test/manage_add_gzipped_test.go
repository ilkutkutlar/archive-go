package archive

import (
  "testing"
  "path"
  "strings"
  "os"
  "fmt"
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
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := createTestFile(tempDir, 1)
  // Remove read permission so adding to archive causes error
  os.Chmod(testFile1, 200)

  gzippedName, actualErr := archive.AddToArchiveGzipped(testFile1, TEST_ARCHIVE, false)
  /* Unlike with tar, since we are not using the -C option, gzip
   * will show the full path of the file, so it won't only say "test.txt" */
  expectedErr := fmt.Sprintf(`Gzip failed:
gzip: %s: Permission denied

exit status 1`, testFile1)

  assertStringEqual(t, gzippedName, "")
  assertStringEqual(t, expectedErr, actualErr.Error())

  // TODO: Should be testing for archiving a directory as well.

  assertFileExists(t, testFile1)
  assertFileDoesNotExist(t, testFile1 + ".gz")

  // Archiving failed - so we expect archive to be empty - i.e. non-existent
  assertFileDoesNotExist(t, TEST_ARCHIVE)

  // TODO: put this in cleanup
  os.Remove(testFile1)
}
