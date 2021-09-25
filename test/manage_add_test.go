package archive

import (
  "testing"
  "os"
  "path"
  "strings"
  archive "example.com/archive/src"
)

func TestAddToArchive(t *testing.T) {
  t.Cleanup(cleanup)

  tempDir := t.TempDir()

  testFile1 := path.Join(tempDir, "test1.txt")
  testDir1 := path.Join(tempDir, "test_dir1")
  testFile2 := path.Join(testDir1, "test2.txt")

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
  testFile1 := path.Join(tempDir, "test1.txt")
  testDir1 := path.Join(tempDir, "test_dir1")
  testFile2 := path.Join(testDir1, "test2.txt")

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
  t.Cleanup(cleanup)

  tempDir := t.TempDir()
  testFile1 := path.Join(tempDir, "test1.txt")
  os.Create(testFile1)
  // Remove read permission so adding to archive causes error
  os.Chmod(testFile1, 200)

  actualErr := archive.AddToArchive(testFile1, TEST_ARCHIVE, false)
  expectedErr := `Adding to archive failed:
tar: test1.txt: Cannot open: Permission denied
tar: Exiting with failure status due to previous errors
exit status 2`

  assertStringEqual(t, expectedErr, actualErr.Error())
  assertFileExists(t, testFile1)
  // Archiving failed - so we expect archive to be empty - i.e. non-existent
  // TODO: although even an empty archive should not have been created!
  assertArchiveContentsEqual(t, TEST_ARCHIVE, "")

  // TODO: put this in cleanup
  os.Remove(testFile1)
}
