package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	build   = "???"
	version = "???"
	short   = false
)

var versionUsage = `Print the app version and build info for the current context. Usage: gupi version [options] Options: --short If true, print just the version number. Default false. `

// setup use function
var versionFunc = func(cmd *Command, args []string) {
	if short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

// setup cmd
func NewVersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	//setup sub-commands
	cmd.flags.BoolVar(&short, "short", false, "")

	//setup use message
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}
	return cmd
}
