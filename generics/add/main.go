package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type IntOrString interface {
	constraints.Integer | string
}

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Printf("Sum of %#v = %v\n", numbers, add(numbers))

	letters := []string{"1", "2", "3", "4"}
	fmt.Printf("Sum of %#v = %v\n", letters, add(letters))
}

// add will accept either int slice or string slice
func add[Type IntOrString](items []Type) (result Type) {
	for _, item := range items {
		result = result + item
	}
	return
}
