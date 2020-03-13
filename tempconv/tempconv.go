// Package tempconv defines several methods for converting temperature across different scales
package tempconv

import "fmt"

// Fahrenheit is a float64 internally
type Fahrenheit float64

// Celsius is a float64 internally
type Celsius float64

// Kelvin is a float64 internally
type Kelvin float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
