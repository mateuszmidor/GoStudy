package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func readWithReadAll() {
	data, _ := ioutil.ReadAll(os.Stdin)
	text := string(data)
	for _, line := range strings.Split(text, "\n") {
		fmt.Println(line)
	}
}

func readWithReader() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n') // \n included in line
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

func readWithScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text() // \n not included in line
		fmt.Println(line)
	}
}

func main() {
	readWithReadAll()
	// readWithReader()
	// readWithScanner()
}
