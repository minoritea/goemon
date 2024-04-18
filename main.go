package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

func main() {
	index := slices.Index(os.Args, "--")
	if index < 1 || index == len(os.Args)-1 {
		fmt.Fprintln(os.Stderr, "Usage: goemon -- command args...")
		os.Exit(1)
	}

	commands := os.Args[index+1:]
	cmd := exec.Command(commands[0], commands[1:]...)
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(cmd.Process.Pid)
}
