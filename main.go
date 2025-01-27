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
	// var cmd *Command

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	fmt.Println("I am ", os.Args[1])

	// switch os.Args[1] {
	// case "version":
	// 	cmd = NewVersionCommand()
	// case "add":
	// 	cmd = NewAddCommand()
	// default:
	// 	usageAndExit(fmt.Sprintf("gc: '%s' is not a gc command.\n", os.Args[1]))
	// }

	// cmd.Init(os.Args[2:])
	// cmd.Run()
}
