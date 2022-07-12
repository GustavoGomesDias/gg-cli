package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// func GetUserNameInLinux() string {
// 	ls := exec.Command("ls")
// 	ls.Dir = "/home/"
// 	stdout, err := ls.Output()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return ""
// 	}

// 	// Print the output
// 	user := strings.ReplaceAll(string(stdout), "\n", "")

// 	return user
// }

func RunCommit(msg string) {

	cmd := exec.Command("/bin/sh", "-c", "git add .; git commit -m '"+msg+"';")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func ValidatePush(msg string, local string, branch string) {
	if msg == "" || local == "" || branch == "" {
		fmt.Print("Msgp, local and branch is required")
		os.Exit(1)
	}
}

func RunCommitWithPush(msg string, local string, branch string) {

	ValidatePush(msg, local, branch)

	cmd := exec.Command("/bin/sh", "-c", "git add .; git commit -m '"+msg+"'; git push "+local+" "+branch+";")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func RunCommitWithPushMain(msg string) {
	cmd := exec.Command("/bin/sh", "-c", "git add .; git commit -m '"+msg+"'; git push origin main")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func RunAmmmend() {
	cmd := exec.Command("git", "commit", "--amend")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func RunGitFunction(gitCmd *flag.FlagSet, commit *string, ammend *bool, push *string, local *string, branch *string, main *string) {
	if *commit != "" {
		message := *commit
		fmt.Println(message)
		RunCommit(message)
		return
	}

	if *ammend {
		RunAmmmend()
		return
	}

	if *push != "" {
		RunCommitWithPush(*push, *local, *branch)
		return
	}

	if *main != "" {
		RunCommitWithPushMain(*main)
		return
	}

	if *commit == "" && !*ammend && *push == "" {
		fmt.Println("Use '--msg' for run commit command, '--am' for run rewrite last commit or '--msgp' for commit and push")
	}
}

func HandleGitCommands() {
	gitCmd := flag.NewFlagSet("g", flag.ExitOnError)
	commit := gitCmd.String("msg", "", "Run commit command with message")
	push := gitCmd.String("msgp", "", "Run commit command with message and push")
	local := gitCmd.String("l", "", "Pass local to a push")
	branch := gitCmd.String("br", "", "Pass branch to a push")
	main := gitCmd.String("main", "", "Run commit and push in origin main")
	ammend := gitCmd.Bool("am", false, "Run amend command")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'g' tag for git actions:")
		fmt.Println("	commit message (--msg), amend command (--am), commit with push (--msgp, --l and --br) or commit with push in main (--main).")
		os.Exit(1)
	}

	if os.Args[1] == "g" {
		gitCmd.Parse(os.Args[2:])
		RunGitFunction(gitCmd, commit, ammend, push, local, branch, main)
	}
}
