package tempconv

// CToF test
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC test
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
