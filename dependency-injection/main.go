package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  uint8
}

type address struct {
	Street string
	Number uint8
}

type hospital struct {
	Director person  `inject:"yes"`
	Location address `inject:"yes"`
}

func main() {
	var h hospital
	injector := prepareInjector()
	injector.Inject(&h)

	fmt.Printf("%+v\n", h)
}

func prepareInjector() Injector {
	director := person{"Smith", 63}
	location := address{"Oak street", 12}

	injector := NewInjector()
	injector.Set(reflect.TypeOf(person{}), director)
	injector.Set(reflect.TypeOf(address{}), location)
	return injector
}
