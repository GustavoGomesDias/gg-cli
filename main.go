package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	gitCmd := flag.NewFlagSet("g", flag.ExitOnError)
	commit := gitCmd.String("msg", "", "Run commit command with message")
	ammend := gitCmd.Bool("am", false, "Run amend command")

	if len(os.Args) < 2 {
		fmt.Println("expected commit message (--msg) or amend command (--am)")
		os.Exit(1)
	}

	if os.Args[1] == "g" {
		gitCmd.Parse(os.Args[2:])
		RunGitFunction(gitCmd, commit, ammend)
	}
}
