package archive

import (
	"fmt"
	archive "github.com/ilkutkutlar/archive/src"
	"os"
	"testing"
)

func TestAddToArchiveGzipped(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchiveGzipped(testFile1, TEST_ARCHIVE, false)
	archive.AddToArchiveGzipped(testDir1, TEST_ARCHIVE, false)

	expected := `test1.txt.gz
test_dir1.tar.gz`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileExists(t, testFile1)
	assertFileExists(t, testDir1)
	// The gzipped files are temporary: must be removed after adding.
	assertFileDoesNotExist(t, testFile1+".gz")
	assertFileDoesNotExist(t, testDir1+".tar.gz")
}

func TestAddToArchiveGzippedAndRemove(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

	archive.AddToArchiveGzipped(testFile1, TEST_ARCHIVE, true)
	archive.AddToArchiveGzipped(testDir1, TEST_ARCHIVE, true)

	expected := `test1.txt.gz
test_dir1.tar.gz`

	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
	assertFileDoesNotExist(t, testFile1)
	assertFileDoesNotExist(t, testDir1)
	assertFileDoesNotExist(t, testFile2)
	// The gzipped files are temporary: must be removed after adding.
	assertFileDoesNotExist(t, testFile1+".gz")
	assertFileDoesNotExist(t, testDir1+".tar.gz")
}

func TestErrorHandledCorrectlyDuringArchivingGzipped(t *testing.T) {
	t.Cleanup(cleanup)
	createTestFiles(t)

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
	assertFileDoesNotExist(t, testFile1+".gz")

	// Archiving failed - so we expect archive to be empty - i.e. non-existent
	assertFileDoesNotExist(t, TEST_ARCHIVE)
}
