package main

import (
	"flag"
	"fmt"
	"os"
)

var usage = `Usage: gc command [options] A simple tool to generate and manage custom templates Options: Commands: add Adds a template to the collection from a local file edit Uses the default text editor to modify a stored template list Lists all stored templates create Generates an instance of a template in the current directory delete Removes a stored template version Prints version info to console `

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}
	flag.Usage()
	os.Exit(0)
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	usageAndExit("")

	var cmd *Command

	switch os.Args[1] {
	case "version":
		// cmd = newVersionCommand()
	default:
		usageAndExit(fmt.Sprintf("gc: '%s' is not a gc command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
