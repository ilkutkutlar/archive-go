package archive

import flag "github.com/spf13/pflag"

import (
	"fmt"
	"os"
	"path"
)

func printHelp() {
	printVersion()
	fmt.Println()
	flag.PrintDefaults()
}

func printVersion() {
	programName := path.Base(os.Args[0])
	versionText := fmt.Sprintf(
		`v0.0.1 - September 2021

Usage: %s {-a|-u|-z} file [-d]
       %s {-h|-l|-t|-v}`,
		programName,
		programName)
	fmt.Println(versionText)
}
