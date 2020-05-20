package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	return &person{name: name, age: 42}
}

func main() {
	fmt.Println(person{"BoB", 20})
	fmt.Println(person{name: "Alice", age: 22})
	fmt.Println(person{name: "Fred"})
	fmt.Println(newPerson("Jon"))

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 55
	fmt.Println(sp.age)
}
