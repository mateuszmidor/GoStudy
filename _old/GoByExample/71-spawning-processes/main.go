package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// simple command, only output
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output() // run command
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// command with input and output
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start() // start the command
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait() // wait for command to exit
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// bash
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
