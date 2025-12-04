// Exercise 2.2: Write a general-purpose unit-conversion program analogous
// to cf that reads numbers from its command-line arguments or from the
// standard input if there are no arguments, and converts each number into
// units like temperature in Celsius and Fahreneit, length in feet and meters,
// weight in pounds and kilograms, and the like.

package converter

import "fmt"

// Temperature types and conversions
type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
