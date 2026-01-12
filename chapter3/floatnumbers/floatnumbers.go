package main

import (
	"fmt"
	"math"
)

func main() {
	// the powers of e with three decimal digits of precision
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eˣ = %8.3f\n", x, math.Exp(float64(x)))
	}
}
