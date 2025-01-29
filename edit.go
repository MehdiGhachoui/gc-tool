package main

import (
	"flag"
	"fmt"
	"os"
)

var editUsage = ``

var editFunc = func(cmd *Command, args []string) {
	fmt.Println("Edit Functionnality")
	os.Exit(0)
}

func NewEditCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("edit", flag.ExitOnError),
		Execute: editFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, editUsage)
	}

	return cmd
}
