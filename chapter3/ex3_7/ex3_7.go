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
	"math/cmplx"
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

// newton returns the color for a point in the complex plane
func newton(z complex128, roots []complex128, rootColors []color.RGBA, maxIter int, tolerance float64) (r, g, b uint8) {
	originalZ := z

	for iter := 0; iter < maxIter; iter++ {
		// Newton's method iteration: z = z - f(z) / f'(z)
		// for f(z) = z^4 - 1, f'(z) = 4z^3
		fz := z*z*z*z - 1
		dfz := 4 * z * z * z

		if cmplx.Abs(dfz) < 1e-12 {
			break // avoid division by zero
		}

		z = z - fz/dfz

		// check convergence to any root
		for i, root := range roots {
			if cmplx.Abs(z-root) < tolerance {
				// color based on root and number of iterations
				return shadeColor(rootColors[i], iter, maxIter, originalZ)
			}
		}
	}

	// black for points that don't converge
	return 0, 0, 0
}

// shadeColor creates a gradient based on iteration count and position
func shadeColor(base color.RGBA, iter, maxIter int, z complex128) (r, g, b uint8) {
	// smooth gradient from dark to light based on iterations
	t := float64(iter) / float64(maxIter)

	// Three different shading effects combined:
	// 1. Brightness based on iterations (faster convergence = brighter)
	brightness := 0.4 + 0.6*(1-t)

	// 2. Subtle texture based on initial position
	texture := 0.95 + 0.05*real(z)*imag(z)

	// 3. Slight hue variation for visual interest
	hueShift := 0.1 * t

	return uint8(float64(base.R) * brightness * texture * (1 - hueShift)),
		uint8(float64(base.G) * brightness * texture),
		uint8(float64(base.B) * brightness * texture * (1 - hueShift))
}
