// Exercise 3.8: Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using
// four different representations of numbers: complex64, complex128, big.Float, and big.Rat. (The latter two types are
// found in the math/big package. Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
// rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts
// become visible?

// Mandelbrot emits a PNG of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"os"
)

func main() {
	// Implementation
}

// 1. complex128 (from book)
func render128() *image.RGBA {
	const size = 256
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cx := float64(x)/size*4 - 2
			cy := float64(y)/size*4 - 2
			img.Set(x, y, mandel128(complex(cx, cy)))
		}
	}

	return img
}

func mandel128(c complex128) color.Color {
	var z complex128
	for n := 0; n < 100; n++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return color.Gray{255 - uint8(n*2)}
		}
	}

	return color.Black
}

// 2. complex64
func render64() *image.RGBA {
	const size = 256
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cx := float32(x)/size*4 - 2
			cy := float32(y)/size*4 - 2
			img.Set(x, y, mandel64(complex(cx, cy)))
		}
	}
	return img
}

func mandel64(c complex64) color.Color {
	var z complex64
	for n := 0; n < 100; n++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return color.Gray{255 - uint8(n*2)}
		}
	}
	return color.Black
}

// 3. big.Float
func renderFloat() *image.RGBA {
	const size = 256
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cx := big.NewFloat(float64(x)/size*4 - 2)
			cy := big.NewFloat(float64(y)/size*4 - 2)
			img.Set(x, y, mandelFloat(cx, cy))
		}
	}

	return img
}

func mandelFloat(cx, cy *big.Float) color.Color {
	zx, zy := new(big.Float), new(big.Float)

	for n := 0; n < 100; n++ {
		// z = z*z + c
		x2 := new(big.Float).Mul(zx, zx)
		y2 := new(big.Float).Mul(zy, zy)
		xy2 := new(big.Float).Mul(zx, zy)
		xy2.Add(xy2, xy2)

		// check if escaped
		r2 := new(big.Float).Add(x2, y2)
		if r2.Cmp(big.NewFloat(4)) > 0 {
			return color.Gray{255 - uint8(n*2)}
		}

		zx.Sub(x2, y2).Add(zx, cx)
		zy.Add(xy2, cy)
	}

	return color.Black
}

// 4. big.Rat
func renderRat() *image.RGBA {
	const size = 256
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cx := big.NewRat(int64(x)*4-int64(size)*2, int64(size))
			cy := big.NewRat(int64(y)*4-int64(size)*2, int64(size))
			img.Set(x, y, mandelRat(cx, cy))
		}
	}

	return img
}

func mandelRat(cx, cy *big.Rat) color.Color {
	zx, zy := new(big.Rat), new(big.Rat)

	for n := 0; n < 100; n++ {
		// z = z*z + c
		x2 := new(big.Rat).Mul(zx, zx)
		y2 := new(big.Rat).Mul(zy, zy)
		xy2 := new(big.Rat).Mul(zx, zy)
		xy2.Add(xy2, xy2)

		// check if escaped
		r2 := new(big.Rat).Add(x2, y2)
		if r2.Cmp(big.NewRat(4, 1)) > 0 {
			return color.Gray{255 - uint8(n*2)}
		}

		zx.Sub(x2, y2).Add(zx, cx)
		zy.Add(xy2, cy)
	}

	return color.Black
}

func save(filnename string, img image.Image) {
	f, err := os.Create(filnename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", f)
	}

	defer f.Close()
	png.Encode(f, img)
}
