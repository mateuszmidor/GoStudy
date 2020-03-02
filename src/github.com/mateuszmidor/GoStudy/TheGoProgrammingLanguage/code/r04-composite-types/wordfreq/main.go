// Project: Count words from stdin
// Usage: go run . < main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	freq := Wordfreq(os.Stdin)

	for word, count := range freq {
		fmt.Printf("%20s - %d\n", word, count)
	}
}
