package main

import "fmt"

type Street struct {
	Name   string
	Number int
}

type Address struct {
	Street   Street
	PostCode string
}

var address = Address{
	PostCode: "33-499",
	Street: Street{
		Name:   "Klonowa",
		Number: 256,
	},
}

const path = "Street.Number"
const expectedValue = 256

func main() {
	fmt.Printf("%+v\n", address)
	fmt.Printf("%s == %v: %t\n", path, expectedValue, FieldAtPathMatches(address, path, expectedValue))
}
