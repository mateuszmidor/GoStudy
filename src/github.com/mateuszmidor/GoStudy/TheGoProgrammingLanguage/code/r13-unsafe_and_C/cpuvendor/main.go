// Project: show how to call C++ function from GO
// Usage: go run .
package main

import "fmt"

func main() {
	fmt.Printf("CPU vendor: %s\n", getCPUVendor())
}
