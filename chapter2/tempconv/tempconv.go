// Package tempconv performs Celsius and Fahrenheit conversions.

// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures
// in the Kelvin scale, where zero Kelvin is -273.15°C and a difference of 1K has the same magnitude as 1°C.
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
