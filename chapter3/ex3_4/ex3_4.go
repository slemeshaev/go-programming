// Exercise 3.4: Following the approach of the Lissajous example in Section 1.7,
// construct a web server that computes surfaces and writes SVG data to the client.
// The server must set the Content-Type header like this:

// w.Header().Set("Content-Type", "image/svg+xml")

// (This step was not required in the Lissjous example because the server
// uses standard heuristics to recommon formats like PNG from the first 512 bytes
// of the response and generates the proper header.) Allow the client to specify
// values like height, width, and color as HTTP request parameters.

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

			// calculate the average height of the polygon's corner
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			// get heights at the four corners of the polygon
			z1 := f(x, y)
			z2 := f(x+xyrange/cells, y)
			z3 := f(x, y+xyrange/cells)
			z4 := f(x+xyrange/cells, y+xyrange/cells)

			// average height of this cell
			avgHeight := (z1 + z2 + z3 + z4) / 4.0

			// generate color based on height
			color := heightToColor(avgHeight)

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Println("</svg>")
}

func heightToColor(z float64) string {
	minZ, maxZ := -0.2, 1.0

	// clamp and normalize
	normalized := (z - minZ) / (maxZ - minZ)

	if normalized < 0 {
		normalized = 0
	}

	if normalized > 1 {
		normalized = 1
	}

	// interpolate between blue (0, 0, 255) and red (255, 0, 0)
	// blue when normalized = 0, red when normalized = 1
	red := int(255 * normalized)
	blue := int(255 * (1 - normalized))

	// format as hex color
	return fmt.Sprintf("#%02x00%02x", red, blue)
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
	// Multiple Peaks
	return math.Sin(x/2)*math.Cos(y/2) + 0.5*math.Sin(x/5)*math.Cos(y/5)
}
