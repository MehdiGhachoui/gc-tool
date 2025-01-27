package main

import (
	"flag"
	"fmt"
	"os"
)

var addUsage = ``

var addFunc = func(cmd *Command, args []string) {
	fmt.Println("Add Functionnality")
	os.Exit(0)
}

func NewAddCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}
