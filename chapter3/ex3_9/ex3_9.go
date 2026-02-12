// Exercise 3.9: Write a web server that renders fractals and writes the
// image data to the client. Allow the client to specify the x, y, and
// zoom values as parameters to the HTTP request.

package main

import (
	"image"
	"image/color"
	"math/cmplx"
)

func main() {
	//
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
