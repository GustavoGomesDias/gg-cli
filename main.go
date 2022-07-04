package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	gitCmd := flag.NewFlagSet("g", flag.ExitOnError)
	commitMsg := gitCmd.String("msg", "", "Commit message")

	if len(os.Args) < 2 {
		fmt.Println("expected commit message (-msg)")
		os.Exit(1)
	}

	if os.Args[1] == "g" {
		gitCmd.Parse(os.Args[2:])
		message := *commitMsg
		fmt.Println(message)
		RunCommit(message)
	}
}
