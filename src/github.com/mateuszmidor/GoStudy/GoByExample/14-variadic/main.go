package main

import "fmt"

func sum(vals ...int) {
	fmt.Print(vals, " ")
	total := 0
	for _, v := range vals {
		total += v
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4)
	vals := []int{5, 10, 15}
	sum(vals...)
}
