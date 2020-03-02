package tempconv

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
