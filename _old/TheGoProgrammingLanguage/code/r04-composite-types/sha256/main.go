// Project: Calc and print sha256 hash
// Usage: go run .
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("%t\n", c1 == c2)
	fmt.Printf("%d\n", calcDiffBitsInSlices(c1[:], c2[:]))
	fmt.Printf("%T\n", c1)
}
