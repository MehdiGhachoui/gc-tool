package main

import (
	"flag"
	"fmt"
	"os"
)

var deleteUsage = ``

var deleteFunc = func(cmd *Command, args []string) {
	fmt.Println("Delete Functionnality")
	os.Exit(0)
}

func NewDeleteCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: deleteFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}

	return cmd
}
