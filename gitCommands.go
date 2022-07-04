package main

import (
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
	// commit := "git add .; git commit -m " + msg

	cmd := exec.Command("/bin/sh", "-c", "git add .; git commit -m '"+msg+"';")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}
