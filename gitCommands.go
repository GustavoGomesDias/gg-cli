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
		fmt.Print("msgp, local and branch is required")
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

	if *ammend != false {
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

	if *commit == "" && *ammend == false && *push == "" {
		fmt.Println("Use '--msg' for run commit command, '--am' for run rewrite last commit or '--msgp' for commit and push")
	}
}
