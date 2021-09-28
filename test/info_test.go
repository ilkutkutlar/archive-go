package archive

import (
	archive "github.com/ilkutkutlar/archive/src"
	"testing"
)

func TestListFilesInArchive(t *testing.T) {
	actual, _ := archive.ListArchive(dummyArchive1)
	expected := `test_dir/
test_dir/test1.txt
test_dir/test_subdir/
test_dir/test_subdir/test2.txt
test.txt
`

	assertStringEqual(t, expected, actual)
}

func TestListFilesInArchiveWithSpacesInName(t *testing.T) {
	actual, _ := archive.ListArchive(dummyArchive2)
	expected := `test.txt
dir with spaces/
dir with spaces/file with spaces.txt
`

	assertStringEqual(t, expected, actual)
}

func TestListTopLevelFilesInArchive(t *testing.T) {
	actual, _ := archive.ListArchiveTopLevel(dummyArchive1)
	expected := `test_dir/
test.txt
`

	assertStringEqual(t, expected, actual)
}

func TestListTopLevelFilesInArchiveWithSpacesInName(t *testing.T) {
	actual, _ := archive.ListArchiveTopLevel(dummyArchive2)
	expected := `test.txt
dir with spaces/
`

	assertStringEqual(t, expected, actual)
}
