package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") // defers are not called when os.Exit is used

	os.Exit(3)
}
