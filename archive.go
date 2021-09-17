package main

import flag "github.com/spf13/pflag"

import (
  "fmt"
  "os"
  "path"
)

var (
  flagAdd = flag.StringP("add", "a", "", "Add file to archive of current directory")
  flagGzipped = flag.StringP("add-gzipped", "z", "", "Add a gzipped version of file to archive. Original file is not affected unless -d is passed")
  flagUnarchive = flag.StringP("unarchive", "u", "", "Unarchive file from archive of current directory")
  flagDelete= flag.BoolP("delete", "d", false, "Pass flag to -a, -u or -z to delete file in dir/archive after operation")
  flagList = flag.BoolP("list", "l", false, "List the files in current directory archive")
  flagTopLevel = flag.BoolP("top-level", "t", false, "List only top-level files and directories in current directory archive")
  flagHelp = flag.BoolP("help", "h", false, "Print this help and exit")
  flagVersion = flag.BoolP("version", "v", false, "Print version and exit")
)

func main() {
  const ARCHIVE_NAME = ".archive.tar"
  flag.CommandLine.SortFlags = false
  flag.Parse()

  if *flagAdd != "" {
    AddToArchive(*flagAdd, ARCHIVE_NAME, *flagDelete)
  } else if *flagGzipped != "" {
    AddToArchiveGzipped(*flagGzipped, ARCHIVE_NAME, *flagDelete)
  } else if *flagUnarchive != "" {
    Unarchive(*flagUnarchive, ARCHIVE_NAME, *flagDelete)
  } else if *flagList {
    ListArchive(ARCHIVE_NAME)
  } else if *flagTopLevel {
    ListArchiveTopLevel(ARCHIVE_NAME)
  } else if *flagHelp {
    printHelp()
  } else if *flagVersion {
    printVersion()
  }
}

func printHelp() {
  printVersion()
  fmt.Println()
  flag.PrintDefaults()
}

func printVersion() {
  programName := path.Base(os.Args[0])
  versionText := fmt.Sprintf(
`v0.0.5 - September 2020

Usage: %s {-a|-u|-z} file [-d]
       %s {-h|-l|-t|-v}`,
    programName,
    programName)
  fmt.Println(versionText)
}
