package main

import (
	"flag"
	"fmt"
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

func RunAmmmend() {
	cmd := exec.Command("git", "commit", "--ammend")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}

func RunGitFunction(gitCmd *flag.FlagSet, commit *string, ammend *bool) {
	if *commit != "" {
		message := *commit
		fmt.Println(message)
		RunCommit(message)
	}

	if *ammend != false {
		RunAmmmend()
	}

	fmt.Println("Use --msg for run commit command or --am for run rewrite last commit")
}
