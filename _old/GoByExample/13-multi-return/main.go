package main

import "fmt"

func vals() (int, int) {
	return 4, 9
}

func main() {
	a, b := vals()
	fmt.Println("a, b =", a, b)

	_, c := vals()
	fmt.Println("c = ", c)
}
