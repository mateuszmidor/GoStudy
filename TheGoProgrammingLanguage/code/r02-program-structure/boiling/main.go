// Project: Display boiling temperature of water in C and F
// Usage: go run .
package main

import "fmt"

const boilingF = 212.0 // package-level visibility

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Water boiling temperature = %gºF or %gºC\n", f, c)
}
