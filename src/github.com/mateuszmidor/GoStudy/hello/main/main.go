package main

import ("fmt"
		"github.com/mateuszmidor/GoStudy/hello/string"
)

func main() {
	msg := "Hello, 世界"
	fmt.Println(msg)
	fmt.Println(string.Reverse(msg))
}
