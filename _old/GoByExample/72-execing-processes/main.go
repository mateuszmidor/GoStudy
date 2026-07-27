package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, loookErr := exec.LookPath("ls")
	if loookErr != nil {
		panic(loookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	// replace current go program with ls
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
