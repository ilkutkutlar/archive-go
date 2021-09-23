package archive

import flag "github.com/spf13/pflag"

var (
  flagAdd = flag.StringP("add", "a", "", "Add file to archive of current directory")
  flagAddGzipped = flag.StringP("add-gzipped", "z", "", "Add a gzipped version of file to archive. Original file is not affected unless -d is passed")
  flagUnarchive = flag.StringP("unarchive", "u", "", "Unarchive file from archive of current directory")
  flagDelete= flag.BoolP("delete", "d", false, "Pass flag to -a, -u or -z to delete file in dir/archive after operation")
  flagList = flag.BoolP("list", "l", false, "List the files in current directory archive")
  flagTopLevel = flag.BoolP("top-level", "t", false, "List only top-level files and directories in current directory archive")
  flagHelp = flag.BoolP("help", "h", false, "Print this help and exit")
  flagVersion = flag.BoolP("version", "v", false, "Print version and exit")
)

const ARCHIVE_NAME = ".archive.tar"

func ParseOptions() {
  flag.CommandLine.SortFlags = false
  flag.Parse()

  if *flagAdd != "" {
    execOptionAdd()
  } else if *flagAddGzipped != "" {
    execOptionAddGzipped()
  } else if *flagUnarchive != "" {
    execOptionUnarchive()
  } else if *flagList {
    execOptionList()
  } else if *flagTopLevel {
    execOptionTopLevel()
  } else if *flagHelp {
    execOptionHelp()
  } else if *flagVersion {
    execOptionVersion()
  }
}

func execOptionAdd() {
  AddToArchive(*flagAdd, ARCHIVE_NAME, *flagDelete)
}

func execOptionAddGzipped() {
  AddToArchiveGzipped(*flagAddGzipped, ARCHIVE_NAME, *flagDelete)
}

func execOptionUnarchive() {
  Unarchive(*flagUnarchive, ARCHIVE_NAME, *flagDelete)
}

func execOptionList() {
  ListArchive(ARCHIVE_NAME)
}

func execOptionTopLevel() {
  ListArchiveTopLevel(ARCHIVE_NAME)
}

func execOptionHelp() {
  printHelp()
}

func execOptionVersion() {
  printVersion()
}
