// Project: implement interface io.Writer to count words written to counter object
// Usage:  go run .
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type counter int

func (c *counter) Write(b []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(b))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(b), nil
}

func main() {
	var c counter
	c.Write([]byte("Litwo Ojczyzno Moja"))
	fmt.Printf("Counted words: %d", c)
}
