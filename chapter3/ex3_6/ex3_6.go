// Exercise 3.6: Supersampling is a technique to reduce
// the effect of pixelation by computing the color value at several
// points within each pixel ant taking the average. The simplest
// method is to divide each pixel into four "subpixels". Implement it.

// Mandelbrot emits a PNG of the Mandelbrot fractal with supersampling.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var colors []color.Color

			for subpy := 0; subpy < 2; subpy++ {
				for subpx := 0; subpx < 2; subpx++ {
					sx := float64(px) + float64(subpx)/2
					sy := float64(py) + float64(subpy)/2

					x := sx/width*(xmax-xmin) + xmin
					y := sy/height*(ymax-уmin) + ymin

					z := complex(x, y)
					colors = append(colors, mandelbrot(z))
				}
			}

			img.Set(px, py, averageColor(colors))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{255, 0, 0, 255}
			case n > 30:
				return color.RGBA{0, 255, 0, 255}
			default:
				return color.RGBA{0, 0, 255, 255}
			}
		}
	}

	return color.Black
}
