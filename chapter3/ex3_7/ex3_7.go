// Exercise 3.7: Another simple fractal uses Newton's method to find
// complex solutions to a function such as z^4-1=0. Shade each starting
// point by the number of iterations required to get close to one of the
// four roots. Color each point by the root it approaches.

// Newton fractal for z^4 - 1 = 0
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
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

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// supersampling: average colors from subpixels
			var r, g, b uint32

			for sy := 0; sy < supersample; sy++ {
				for sx := 0; sx < supersample; sx++ {
					// subpixel coordinates with jitter for better sampling
					sxPos := float64(px) + (float64(sx)+0.5)/float64(supersample)
					syPos := float64(py) + (float64(sy)+0.5)/float64(supersample)

					// map to complex plane
					x := sxPos/float64(width)*(xmax-xmin) + xmin
					y := syPos/float64(height)*(ymax-ymin) + ymin

					// compute color for this subpixel
					cr, cg, cb := newton(complex(x, y), roots, rootColors, maxIter, tolerance)

					r += uint32(cr)
					g += uint32(cg)
					b += uint32(cb)
				}
			}

			// average the colors
			samples := supersample * supersample
			img.Set(px, py, color.RGBA{
				uint8(r / uint32(samples)),
				uint8(g / uint32(samples)),
				uint8(b / uint32(samples)),
				255,
			})
		}
	}

	png.Encode(os.Stdout, img)
}
