// Project: sum numbers using variadic function
// Usage: go run .
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}

func sum(numbers ...int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}

	return result
}
