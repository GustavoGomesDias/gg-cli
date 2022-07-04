package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	gitCmd := flag.NewFlagSet("g", flag.ExitOnError)
	commit := gitCmd.String("msg", "", "Run commit command with message")
	push := gitCmd.String("msgp", "", "Run commit command with message and push")
	local := gitCmd.String("l", "", "Pass local to a push")
	branch := gitCmd.String("br", "", "Pass branch to a push")
	ammend := gitCmd.Bool("am", false, "Run amend command")

	if len(os.Args) < 2 {
		fmt.Println("expected commit message (--msg), amend command (--am) or commit with push (--msgp, --l and --br)")
		os.Exit(1)
	}

	if os.Args[1] == "g" {
		gitCmd.Parse(os.Args[2:])
		RunGitFunction(gitCmd, commit, ammend, push, local, branch)
	}
}
