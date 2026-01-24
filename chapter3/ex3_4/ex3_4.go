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
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth, defaultHeight = 600, 320    // canvas size in pixels
	defaultCells                = 100         // number of grid cells
	defaultXyrange              = 30.0        // axis ranges (-xyrange..+xyrange)
	defaultAngle                = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(defaultAngle), math.Cos(defaultAngle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// parse query parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get parameters or use defaults
	width := getIntParam(r, "width", defaultWidth)
	height := getIntParam(r, "height", defaultHeight)
	cells := getIntParam(r, "cells", defaultCells)
	xyrange := getFloatParam(r, "xyrange", defaultXyrange)
	color := r.FormValue("color")
	if color == "" {
		color = "height" // default to height-based coloring
	}

	// calculate derived constants
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	// set the Content-Type header
	w.Header().Set("Content-Type", "image/svg+xml")

	// generate svg
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j, width, height, xyrange, xyscale, zscale, cells)
			bx, by, ok2 := corner(i, j, width, height, xyrange, xyscale, zscale, cells)
			cx, cy, ok3 := corner(i, j+1, width, height, xyrange, xyscale, zscale, cells)
			dx, dy, ok4 := corner(i+1, j+1, width, height, xyrange, xyscale, zscale, cells)

			// if at least one vertex of a polygon is invalid, we skip the poligon.
			if !(ok1 && ok2 && ok3 && ok4) {
				continue
			}

			// calculate the average height of the polygon's corner
			x := xyrange * (float64(i)/float64(cells) - 0.5)
			y := xyrange * (float64(j)/float64(cells) - 0.5)

			// get heights at the four corners of the polygon
			z1 := f(x, y)
			z2 := f(x+xyrange/float64(cells), y)
			z3 := f(x, y+xyrange/float64(cells))
			z4 := f(x+xyrange/float64(cells), y+xyrange/float64(cells))

			// average height of this cell
			avgHeight := (z1 + z2 + z3 + z4) / 4.0

			// determine fill color based on parameter
			fillColor := getFillColor(color, avgHeight)

			fmt.Fprintf(w, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, width, height int, xyrange, xyscale, zscale float64, cells int) (float64, float64, bool) {
	// find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// compute surface height z
	z := f(x, y)

	// check that z is a finite number
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// project (x, y, z) isometrically onto 2D SVG canvas (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

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

func getFillColor(colorScheme string, height float64) string {
	switch colorScheme {
	case "red":
		return "#ff0000"
	case "blue":
		return "#0000ff"
	case "green":
		return "#00ff00"
	case "yellow":
		return "#ffff00"
	case "purple":
		return "#800080"
	case "orange":
		return "#ffa500"
	case "height":
		// height-based coloring (blue to red gradient)
		minZ, maxZ := -0.2, 1.0
		normalized := (height - minZ) / (maxZ - minZ)

		if normalized < 0 {
			normalized = 0
		}

		if normalized > 1 {
			normalized = 1
		}

		red := int(255 * normalized)
		blue := int(255 * (1 - normalized))
		return fmt.Sprintf("#%02x00%02x", red, blue)
	default:
		// if color is a hex code, use it directly
		if len(colorScheme) == 7 && colorScheme[0] == '#' {
			return colorScheme
		}

		// default to white
		return "#ffffff"
	}
}

func getIntParam(r *http.Request, name string, defaultValue int) int {
	value := r.FormValue(name)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil || intValue <= 0 {
		return defaultValue
	}

	return intValue
}

func getFloatParam(r *http.Request, name string, defaultValue float64) float64 {
	value := r.FormValue(name)
	if value == "" {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil || floatValue <= 0 {
		return defaultValue
	}

	return floatValue
}
