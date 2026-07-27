package main

import "fmt"

type Vector[T any] struct {
	Items []T
}

func (v *Vector[T]) Append(item T) *Vector[T] {
	v.Items = append(v.Items, item)
	return v
}

func main() {
	ints := &Vector[int]{}
	ints.Append(1).Append(2).Append(3).Append(4)
	fmt.Printf("%#v\n", ints.Items)

	strings := &Vector[string]{}
	strings.Append("1").Append("2").Append("3").Append("4")
	fmt.Printf("%#v\n", strings.Items)
}
