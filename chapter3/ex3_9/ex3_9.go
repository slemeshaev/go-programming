// Exercise 3.9: Write a web server that renders fractals and writes the
// image data to the client. Allow the client to specify the x, y, and
// zoom values as parameters to the HTTP request.

package main

import (
	"image"
	"image/color"
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
	// implementation
	return color.Black
}
