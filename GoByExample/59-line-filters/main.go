package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("/proc/cpuinfo")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
