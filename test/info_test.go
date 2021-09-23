package archive

import (
  "testing"
  archive "example.com/archive/src"
)

func TestListFilesInArchive(t *testing.T) {
  actual, _ := archive.ListArchive("fixtures/dummy_archive.tar")
  expected := `test_dir/
test_dir/test1.txt
test_dir/test_subdir/
test_dir/test_subdir/test2.txt
test.txt
`

  if actual != expected {
    t.Fail()
  }
}

func TestListFilesInArchiveWithSpacesInName(t *testing.T) {
  actual, _ := archive.ListArchive("fixtures/dummy_archive2.tar")
  expected := `test.txt
dir with spaces/
dir with spaces/file with spaces.txt
`

  if actual != expected {
    t.Fail()
  }
}


func TestListTopLevelFilesInArchive(t *testing.T) {
  actual, _ := archive.ListArchiveTopLevel("fixtures/dummy_archive.tar")
  expected := `test_dir/
test.txt
`

  if actual != expected {
    t.Fail()
  }
}


func TestListTopLevelFilesInArchiveWithSpacesInName(t *testing.T) {
  actual, _ := archive.ListArchiveTopLevel("fixtures/dummy_archive2.tar")
  expected := `test.txt
dir with spaces/
`

  if actual != expected {
    t.Fail()
  }
}
