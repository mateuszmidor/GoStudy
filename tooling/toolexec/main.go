package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("must receive tool and args, received: %v\n", os.Args)
		os.Exit(1)
	}
	tool := os.Args[1]
	args := os.Args[2:]

	fmt.Fprintln(os.Stderr, "TOOLEXEC:", tool)
	for _, arg := range args {
		fmt.Fprintln(os.Stderr, " ", arg)
	}

	cmd := exec.Command(tool, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
