// Exercise 3.7: Another simple fractal uses Newton's method to find
// complex solutions to a function such as z^4-1=0. Shade each starting
// point by the number of iterations required to get close to one of the
// four roots. Color each point by the root it approaches.

// Newton fractal for z^4 - 1 = 0
package main

import (
	"fmt"
	"image/color"
)

func main() {
	const (
		width, height          = 1024, 1024
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		maxIter                = 40
		tolerance              = 1e-6
		supersample            = 2 // 2x2 supersampling
	)

	// roots of z^4 - 1 = 0
	roots := []complex128{
		1 + 0i,
		0 + 1i,
		-1 + 0i,
		0 - 1i,
	}

	// vibrant colors for each root
	rootColors := []color.RGBA{
		{255, 50, 50, 255},  // red for 1
		{50, 255, 50, 255},  // green for i
		{80, 80, 255, 255},  // blue for -1
		{255, 255, 50, 255}, // yellow for -i
	}

	fmt.Println(roots, rootColors)
}
