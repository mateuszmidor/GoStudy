// Project: return a value from function that returns nothing
// Usage: go run .
package main

import "fmt"

func main() {
	defer recoverPanic()
	doPanic()
}

func doPanic() {
	panic(42)
}

func recoverPanic() {
	p := recover()
	fmt.Printf("Recovered from panic of value: %v\n", p)
}
