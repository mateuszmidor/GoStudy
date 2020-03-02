// Project: show how a package can be auto initialized
// Usage: go run .
package main

import (
	_ "autoinit" // autoinit will be included and its "init()" will be run before main()
	"fmt"
)

func main() {
	fmt.Println("Hello from main!")
}
