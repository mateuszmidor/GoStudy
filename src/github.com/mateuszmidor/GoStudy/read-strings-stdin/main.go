package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text() // \n not included in line
		fmt.Println(line)
	}
}

func main() {
	// readWithReader()
	readWithScanner()
}
