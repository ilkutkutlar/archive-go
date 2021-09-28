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
	copyFile(dummyArchive1, testArchive)

	err := archive.Unarchive("test_dir", testArchive, false)

	assertFileExists(t, path.Join(path.Dir(testArchive), "test_dir"))
	assertFileExists(t, path.Join(path.Dir(testArchive), "test_dir", "test1.txt"))
	assertNil(t, err)

	expected := `test_dir/
test_dir/test1.txt
test_dir/test_subdir/
test_dir/test_subdir/test2.txt
test.txt`
	assertArchiveContentsEqual(t, testArchive, expected)
}

func TestUnarchiveFilesWithSpacesInNameFromArchive(t *testing.T) {
	t.Cleanup(cleanup)
	/* dummy archive 2 contains:
	 * test.txt, dir with spaces/, dir with spaces/file with spaces.txt */
	copyFile(dummyArchive2, testArchive)

	err := archive.Unarchive("dir with spaces", testArchive, false)

	assertFileExists(t, path.Join(path.Dir(testArchive), "dir with spaces"))
	assertFileExists(t, path.Join(path.Dir(testArchive), "dir with spaces", "file with space.txt"))
	assertNil(t, err)

	expected := `test.txt
dir with spaces/
dir with spaces/file with spaces.txt`
	assertArchiveContentsEqual(t, testArchive, expected)
}

func TestUnarchiveFilesFromArchiveAndRemove(t *testing.T) {
	t.Cleanup(cleanup)

	/* dummy archive contains:
	 * test.txt, test_dir/, test_dir/test1.txt */
	copyFile(dummyArchive1, testArchive)

	err := archive.Unarchive("test_dir", testArchive, true)

	assertFileExists(t, path.Join(path.Dir(testArchive), "test_dir"))
	assertFileExists(t, path.Join(path.Dir(testArchive), "test_dir", "test1.txt"))
	assertNil(t, err)

	expected := "test.txt"
	assertArchiveContentsEqual(t, testArchive, expected)
}

func TestUnarchiveFilesWithSpacesInNameFromArchiveAndRemove(t *testing.T) {
	t.Cleanup(cleanup)
	/* dummy archive 2 contains:
	 * test.txt, dir with spaces/, dir with spaces/file with spaces.txt */
	copyFile(dummyArchive2, testArchive)

	err := archive.Unarchive("dir with spaces", testArchive, true)

	assertFileExists(t, path.Join(path.Dir(testArchive), "dir with spaces"))
	assertFileExists(t, path.Join(path.Dir(testArchive), "dir with spaces", "file with space.txt"))
	assertNil(t, err)

	expected := "test.txt"
	assertArchiveContentsEqual(t, testArchive, expected)
}

func TestErrorHandledCorrectlyUnarchiving(t *testing.T) {
	t.Skip()
}
