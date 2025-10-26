package main

import "fmt"

// Celsius is temperature in ºcelsius
type Celsius float64

// Fahrenheit is temperatue in ºfarenheit
type Fahrenheit float64

// Kelvin is temperature in Kelvin's units
type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }

// CToF converts Celsius -> Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts Celsius -> Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToC converts Fahrenheit -> Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts Kelvin -> Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
