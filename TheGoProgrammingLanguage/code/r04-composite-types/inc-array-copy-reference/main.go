package main

import (
	"fmt"
)

func inc(t []int) {
	for i := range t {
		t[i]++
	}
}
func main() {
	t := [...]int{1, 2, 3, 4}
	fmt.Println("Original array:\n", t)

	// increment a copy of array
	t2 := make([]int, len(t))
	copy(t2, t[:]) // t[:] creates a slice referencing entire array t. Then copy its contents into separate slice t2
	inc(t2)
	fmt.Println("Increment a copy:\n", t, t2)

	// increment reference of array
	t3 := t[:]
	inc(t3)
	fmt.Println("Increment a reference:\n", t, t3)
}
