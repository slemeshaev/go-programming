// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	blackIndex = 1
)

func main() {
	file, err := os.Create("lissajous.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}

	defer file.Close()

	err = lissajous(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating GIF: %v\n", err)
		return
	}

	fmt.Printf("GIF successfuly created: %s", file.Name())
}

func lissajous(out io.Writer) error {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resulution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	freq := r.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return gif.EncodeAll(out, &anim)
}
