package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if n := 9; n < 0 {
		fmt.Println("n is negative")
	} else if n < 10 {
		fmt.Println("n has 1 digit")
	} else {
		fmt.Println("n has multiple digits")
	}
}
