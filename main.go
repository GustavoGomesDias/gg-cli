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
	main := gitCmd.String("main", "", "Run commit and push in origin main")
	ammend := gitCmd.Bool("am", false, "Run amend command")

	if len(os.Args) < 2 {
		fmt.Println("Expected commit message (--msg), amend command (--am), commit with push (--msgp, --l and --br) or commit with push in main (--main).")
		os.Exit(1)
	}

	if os.Args[1] == "g" {
		gitCmd.Parse(os.Args[2:])
		RunGitFunction(gitCmd, commit, ammend, push, local, branch, main)
	}
}
