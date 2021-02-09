package main

import (
	"fmt"
	"os"
	"reflect"
)

// PrettyPrint prints object in nice nested way
func PrettyPrint(resource interface{}) {
	prettyPrintRecursive(reflect.TypeOf(resource), reflect.ValueOf(resource), 0)
}

func prettyPrintRecursive(t reflect.Type, v reflect.Value, level int) {
	// sentinel for infinite recursion error
	if level > 10 {
		panic("Level > 10. Recursion too deep")
	}

	switch v.Kind() {

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			printfl("%s:\n", level, t.Field(i).Name)
			prettyPrintRecursive(v.Field(i).Type(), v.Field(i), level+1)
		}

	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			printfl("[no items]\n", level)

		} else {
			for i := 0; i < count; i++ {
				printfl("[%d]\n", level, i)
				s := v.Index(i)
				prettyPrintRecursive(s.Type(), s, level)
			}
		}

	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			printfl("[empty]\n", level)
		} else {
			prettyPrintRecursive(v.Elem().Type(), v.Elem(), level)
		}

	default:
		printfl("%v\n", level, v)
	}

}

// Indented Printf
func printfl(format string, level int, args ...interface{}) {
	fmt.Printf("%*s", level*2, "")
	fmt.Printf(format, args...)
}

func main() {
	PrettyPrint(os.Stderr)
}
