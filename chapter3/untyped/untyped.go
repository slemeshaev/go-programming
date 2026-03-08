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
}
