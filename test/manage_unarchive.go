package archive

import (
	archive "github.com/ilkutkutlar/archive/src"
	"path"
	"testing"
)

func TestUnarchiveFilesFromArchive(t *testing.T) {
	t.Cleanup(cleanup)

	/* dummy archive contains:
	 * test.txt, test_dir/, test_dir/test1.txt */
	copyFile(DUMMY_ARCHIVE_1, TEST_ARCHIVE)

	err := archive.Unarchive("test_dir", TEST_ARCHIVE, false)

	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "test_dir"))
	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "test_dir", "test1.txt"))
	assertNil(t, err)

	expected := `test_dir/
test_dir/test1.txt
test_dir/test_subdir/
test_dir/test_subdir/test2.txt
test.txt`
	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
}

func TestUnarchiveFilesWithSpacesInNameFromArchive(t *testing.T) {
	t.Cleanup(cleanup)
	/* dummy archive 2 contains:
	 * test.txt, dir with spaces/, dir with spaces/file with spaces.txt */
	copyFile(DUMMY_ARCHIVE_2, TEST_ARCHIVE)

	err := archive.Unarchive("dir with spaces", TEST_ARCHIVE, false)

	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "dir with spaces"))
	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "dir with spaces", "file with space.txt"))
	assertNil(t, err)

	expected := `test.txt
dir with spaces/
dir with spaces/file with spaces.txt`
	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
}

func TestUnarchiveFilesFromArchiveAndRemove(t *testing.T) {
	t.Cleanup(cleanup)

	/* dummy archive contains:
	 * test.txt, test_dir/, test_dir/test1.txt */
	copyFile(DUMMY_ARCHIVE_1, TEST_ARCHIVE)

	err := archive.Unarchive("test_dir", TEST_ARCHIVE, true)

	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "test_dir"))
	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "test_dir", "test1.txt"))
	assertNil(t, err)

	expected := "test.txt"
	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
}

func TestUnarchiveFilesWithSpacesInNameFromArchiveAndRemove(t *testing.T) {
	t.Cleanup(cleanup)
	/* dummy archive 2 contains:
	 * test.txt, dir with spaces/, dir with spaces/file with spaces.txt */
	copyFile(DUMMY_ARCHIVE_2, TEST_ARCHIVE)

	err := archive.Unarchive("dir with spaces", TEST_ARCHIVE, true)

	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "dir with spaces"))
	assertFileExists(t, path.Join(path.Dir(TEST_ARCHIVE), "dir with spaces", "file with space.txt"))
	assertNil(t, err)

	expected := "test.txt"
	assertArchiveContentsEqual(t, TEST_ARCHIVE, expected)
}

func TestErrorHandledCorrectlyUnarchiving(t *testing.T) {
	t.Skip()
}
