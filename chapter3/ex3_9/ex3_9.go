// Exercise 3.9: Write a web server that renders fractals and writes the
// image data to the client. Allow the client to specify the x, y, and
// zoom values as parameters to the HTTP request.

// Request examples:
// http://localhost:8080
// http://localhost:8080/?x=-0.743643887037151&y=0.13182590420533&zoom=1000
// http://localhost:8080/?x=0.28&y=0.008&zoom=50
// http://localhost:8080/?x=-1.769&y=0&zoom=100

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server starting on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// parse query parameters with defaults
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		log.Printf("Error parsing x parameter %v:", err)
		return
	}

	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		log.Printf("Error parsing y parameter %v:", err)
		return
	}

	zoom, err := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64)
	if err != nil {
		log.Printf("Error parsing zoom parameter %v:", err)
		return
	}

	// set defaults
	if x == 0 && y == 0 {
		x, y = -0.5, 0 // classic Mandelbrot view
	}

	if zoom == 0 {
		zoom = 1
	}

	// calculate bounds
	const size = 1024
	scale := 2.0 / zoom
	xmin, xmax := x-scale, x+scale
	ymin, ymax := y-scale, y+scale

	// render fractal
	img := renderMandelbrot(xmin, xmax, ymin, ymax, size)

	// set content type and send image
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func renderMandelbrot(xmin, xmax, ymin, ymax float64, size int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for py := 0; py < size; py++ {
		y := float64(py)/float64(size)*(ymax-ymin) + ymin
		for px := 0; px < size; px++ {
			x := float64(px)/float64(size)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}
