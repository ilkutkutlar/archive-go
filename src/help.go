package archive

import flag "github.com/spf13/pflag"

import (
	"fmt"
	"os"
	"path"
)

// PrintHelp prints a help text summarising the functions of flags
func PrintHelp() {
	PrintVersion()
	fmt.Println()
	flag.PrintDefaults()
}

// PrintVersion prints the version of archive and a small usage text
func PrintVersion() {
	programName := path.Base(os.Args[0])
	versionText := fmt.Sprintf(
		`v0.0.2 - September 2021

Usage: %s {-a|-u} FILE [-d|-z]
       %s {-h|-l|-t|-v}`,
		programName,
		programName)
	fmt.Println(versionText)
}
