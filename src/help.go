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
		`v0.0.2 - September 2021

Usage: %s {-a|-u} FILE [-d|-z]
       %s {-h|-l|-t|-v}`,
		programName,
		programName)
	fmt.Println(versionText)
}
