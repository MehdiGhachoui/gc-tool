package main

import (
	"flag"
	"fmt"
	"os"
)

var usage = `This is a custom Command line template`

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}
	flag.Usage()
	os.Exit(0)
}

func main() {
	var cmd *Command

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			cmd = NewVersionCommand()
		case "add":
			cmd = NewAddCommand()
		case "edit":
			cmd = NewEditCommand()
		case "delete":
			cmd = NewDeleteCommand()
		default:
			usageAndExit(fmt.Sprintf("gc: '%s' is not a gc command.\n", os.Args[1]))
		}

		cmd.Init(os.Args[2:])
		cmd.Run()
	} else {
		flag.Usage = func() {
			fmt.Fprint(os.Stderr, fmt.Sprint(usage))
		}
		usageAndExit("")
	}

}
