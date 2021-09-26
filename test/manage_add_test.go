package archive

import (
	archive "example.com/archive/src"
	"os"
	"testing"
)

func TestAddToArchive(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchive(testFile1, TEST_ARCHIVE, false)
	archive.AddToArchive(testDir1, TEST_ARCHIVE, false)

	expected := `test1.txt
test_dir1/
test_dir1/test2.txt`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileExists(t, testFile1)
	assertFileExists(t, testDir1)
}

func TestAddToArchiveFilesWithSpaces(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchive(testFile3, TEST_ARCHIVE, false)
	archive.AddToArchive(testDir2, TEST_ARCHIVE, false)

	expected := `test file 1.txt
test dir/
test dir/test file 2.txt`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileExists(t, testFile3)
	assertFileExists(t, testDir2)
}

func TestAddToArchiveAndRemove(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchive(testFile1, TEST_ARCHIVE, true)
	archive.AddToArchive(testDir1, TEST_ARCHIVE, true)

	expected := `test1.txt
test_dir1/
test_dir1/test2.txt`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileDoesNotExist(t, testFile1)
	assertFileDoesNotExist(t, testDir1)
}

func TestAddToArchiveFilesWithSpacesAndRemove(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchive(testFile3, TEST_ARCHIVE, true)
	archive.AddToArchive(testDir2, TEST_ARCHIVE, true)

	expected := `test file 1.txt
test dir/
test dir/test file 2.txt`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileDoesNotExist(t, testFile3)
	assertFileDoesNotExist(t, testDir2)
}

func TestErrorHandledCorrectlyDuringArchiving(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

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
}
