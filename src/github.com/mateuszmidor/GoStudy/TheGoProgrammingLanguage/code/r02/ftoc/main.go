package main

// Project: Fahrenheit to Celcius conversion
// Usage: ./fahrenheit

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF = %gºC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gºF = %gºC\n", boilingF, fToC(boilingF))
}

func fToC(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
