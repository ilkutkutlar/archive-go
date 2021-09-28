package archive

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

const defaultArchiveName = ".archive.tar"

var (
	flagAdd         = flag.StringP("add", "a", "", "Add file to archive of current directory")
	flagUnarchive   = flag.StringP("unarchive", "u", "", "Unarchive file from archive of current directory")
	flagArchiveName = flag.StringP("archive-name", "n", defaultArchiveName, "Use a custom archive name instead of the default .archive.tar")
	flagGzip        = flag.BoolP("gzip", "z", false, "Used with -a to gzip the file/dir before archiving it. Original file is not affected (i.e. not gzipped) but will be deleted if -d is passed.")
	flagDelete      = flag.BoolP("delete", "d", false, "Pass flag to -a, -u or -z to delete file in dir/archive after operation")
	flagList        = flag.BoolP("list", "l", false, "List the files in current directory archive")
	flagTopLevel    = flag.BoolP("top-level", "t", false, "List only top-level files and directories in current directory archive")
	flagHelp        = flag.BoolP("help", "h", false, "Print this help and exit")
	flagVersion     = flag.BoolP("version", "v", false, "Print version and exit")
)

// ParseOptions parses the flags passed and executes the appropriate options.
func ParseOptions() {
	flag.CommandLine.SortFlags = false
	flag.Parse()

	if *flagAdd != "" {
		execOptionAdd()
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
	} else {
		execOptionHelp()
	}
}

func execOptionAdd() {
	if *flagGzip {
		execAdd()
	} else {
		execAddGzipped()
	}
}

func execAdd() {
	gzippedFileName, err := AddToArchiveGzipped(*flagAdd, *flagArchiveName, *flagDelete)

	if err == nil {
		fmt.Println(*flagAdd, "added to archive as a gzipped file with name", gzippedFileName)
	} else {
		fmt.Print(err)
		os.Exit(1)
	}
}

func execAddGzipped() {
	err := AddToArchive(*flagAdd, *flagArchiveName, *flagDelete)

	if err == nil {
		fmt.Println(*flagAdd, "added to archive")
	} else {
		fmt.Print(err)
		os.Exit(1)
	}
}

func execOptionUnarchive() {
	err := Unarchive(*flagUnarchive, *flagArchiveName, *flagDelete)

	if err == nil {
		fmt.Println("Retrieved", *flagUnarchive, "from archive")

		if *flagDelete {
			fmt.Println("Deleted", *flagUnarchive, "from archive permanently")
		}
	} else {
		fmt.Print(err)
		os.Exit(1)
	}
}

func execOptionList() {
	out, err := ListArchive(*flagArchiveName)

	if err == nil {
		fmt.Println("Files in archive:")
		fmt.Print(out)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execOptionTopLevel() {
	out, err := ListArchiveTopLevel(*flagArchiveName)

	if err == nil {
		fmt.Println("Top-level files in archive:")
		fmt.Print(out)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execOptionHelp() {
	PrintHelp()
}

func execOptionVersion() {
	PrintVersion()
}
