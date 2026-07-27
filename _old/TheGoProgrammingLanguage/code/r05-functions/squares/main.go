// Project: subsequent square number generator using closures
// Usage: go run .
package main

import "fmt"

func main() {
	squareGenerator := squares()
	fmt.Println(squareGenerator()) // 1
	fmt.Println(squareGenerator()) // 4
	fmt.Println(squareGenerator()) // 9
	fmt.Println(squareGenerator()) // 16
}

func squares() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}
