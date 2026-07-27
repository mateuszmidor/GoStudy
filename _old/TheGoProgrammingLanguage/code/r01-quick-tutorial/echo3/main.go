// Project: print main arguments separated with spaces. Using strings.Join
// Usage: go run . one two three
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
