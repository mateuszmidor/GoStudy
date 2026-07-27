package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type person struct {
	Name     string
	Age      uint8
	Siblings []string
}

const filename = "/tmp/person.dat"

// error handling helper
func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %v. Exit now\n", err)
		os.Exit(1)
	}
}

func write(p person) {
	file, err := os.Create(filename)
	exitOnError(err)
	defer file.Close()

	enc := gob.NewEncoder(file) // will write to file
	err = enc.Encode(&p)
	exitOnError(err)

	fmt.Printf("Written person:\n %+v\n", p)
}

func read() (p person) {
	file, err := os.Open(filename)
	exitOnError(err)
	defer file.Close()

	dec := gob.NewDecoder(file) // will read from file
	err = dec.Decode(&p)
	exitOnError(err)

	fmt.Printf("Read person:\n %+v\n", p)
	return
}

func main() {
	p1 := person{
		Name:     "Andrzej",
		Age:      35,
		Siblings: []string{"Franko", "Przemo"},
	}

	write(p1)
	_ = read()
}
