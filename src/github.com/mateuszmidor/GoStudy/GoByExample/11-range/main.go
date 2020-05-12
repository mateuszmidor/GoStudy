package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// index, value
	for _, value := range nums {
		sum += value
	}
	fmt.Println("sum:", sum)

	for index, value := range nums {
		if value == 3 {
			fmt.Println("index of 3:", index)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over just keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// byte index, unicode code point
	for i, c := range "półciężarówka" {
		fmt.Printf("%2.2d %3.3[2]d %[2]c\n", i, c)
	}
}
