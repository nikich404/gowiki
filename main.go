package main

import (
	"fmt"
	"gowiki/cli"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: gowiki <build|serve|new> [args]")
		os.Exit(1)
	}
	command := os.Args[1]

	switch command {
	case "build":
		if err := cli.RunBuild(); err != nil {
			fmt.Println("build error:", err)
			os.Exit(1)
		}
	case "serve":
		cli.RunServer()
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("usage: gowiki new <title>")
			os.Exit(1)
		}
		if err := cli.RunNew(os.Args[2]); err != nil {
			fmt.Println("new error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("unknown command:", command)
		fmt.Println("usage: gowiki <build|serve|new> [args]")
		os.Exit(1)
	}

	fmt.Println("done")
}
