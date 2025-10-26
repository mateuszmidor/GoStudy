package main

import "fmt"

func main() {
	t := [...]int{1, 2, 3, 4, 5}  // arrray, not slice
	fmt.Printf("%[1]T, %[1]v", t) // refer to first argument twice
}
