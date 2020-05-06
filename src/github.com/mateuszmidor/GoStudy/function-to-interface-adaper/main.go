package main

import "fmt"

// Mutator returns mutated version of the input string
type Mutator interface {
	mutate(s string) string
}

// MutatorAdapter turns regular function of signature "func(string) string" into a type that implements Mutator interface
type MutatorAdapter func(s string) string

// THIS does the trick: mutate method of MutatorAdapter calls the owning function itself
func (f MutatorAdapter) mutate(s string) string {
	return f(s)
}

// 1. verbose approach - struct implementing the Mutator interface
type structMutator struct {
}

func (m structMutator) mutate(s string) string {
	return s + s
}

// 2. concise approach - just a function with signature of Mutator.mutate
func funcMutator(s string) string {
	return "?" + s + "?"
}

// mutate string using mutator and print it out
func drawMutated(m Mutator, s string) {
	newString := m.mutate(s)
	fmt.Println(newString)
}

func main() {
	drawMutated(structMutator{}, "fafa")             // like http.Handle
	drawMutated(MutatorAdapter(funcMutator), "fafa") // like http.HandlerFunc
}
