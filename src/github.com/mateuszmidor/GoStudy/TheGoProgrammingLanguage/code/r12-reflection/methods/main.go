// Project: print all methods of a type
// Usage: go run .
package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	print(time.Hour)
}

func print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methodType.String(), "func"))
	}
}
