// Exercise 3.2: Experiment with visulizations of other functions from
// the math package. Can you produce an egg box, moguls, or a saddle?

// Surface computes an SVG rendering of a 3D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)

			// if at least one vertex of a polygon is invalid, we skip the poligon.
			if !(ok1 && ok2 && ok3 && ok4) {
				continue
			}

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height z
	z := f(x, y)

	// check that z is a finite number
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	// check that the final coordinates are finite numbers
	if math.IsInf(sx, 0) || math.IsNaN(sx) || math.IsInf(sy, 0) || math.IsNaN(sy) {
		return 0, 0, false
	}

	return sx, sy, true
}

func f(x, y float64) float64 {
	// 1. Egg Box Pattern
	// return 0.2 * (math.Sin(x) + math.Sin(y))
	// 2. Muguls (Sinusoidal hills)
	// return math.Sin(x) * math.Cos(y)
	// 3. Haddle (Hyperbolic parabaloid)
	return (x*x - y*y) * 0.05
}
