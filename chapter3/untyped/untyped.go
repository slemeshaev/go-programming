package main

import (
	"fmt"
	"math"
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	fmt.Printf("x = %g\ny = %g\nz = %g\n", x, y, z)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)
	fmt.Println(5 / 9 * (f - 32))
	fmt.Println(5.0 / 9.0 * (f - 32))

	var g float64 = 3 + 0i // untyped complex -> float64
	g = 2                  // untyped integer -> float64
	g = 1e123              // untyped floating point -> float64
	g = 'a'                // untyped rune -> float64

	fmt.Println(g)

	var h float64 = float64(3 + 0i)
	h = float64(2)
	h = float64(1e123)
	h = float64('a')

	fmt.Println(h)

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}
