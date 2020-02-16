// Project: demonstarte how to register new type of command line flag under "flag" utility
// Usage: go run . -temp 20C
//                 -temp 0F
package main

import (
	"flag"
	"fmt"
)

// celsiusFlag implements flag.Value interface
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no need to check for errors as unrecognized value is checked
	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
	case "F", "ºF":
		f.Celsius = FToC(Fahrenheit(value))
	default:
		return fmt.Errorf("Unrecognized temperature %q", s)
	}
	return nil // OK
}

// CelsiusFlag registers new type of flag under "flag" utility and returns pointer to parsed temperature
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp *Celsius = CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
