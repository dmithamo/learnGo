package tempconv

// CToF converts temp from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts temp from Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

// FToC converts temp from Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts temp from Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

// KToC converts temp from Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

// KToF converts temp from Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
