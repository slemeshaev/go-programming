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

const (
	whiteIndex  = 0
	blackIndex  = 1
	redIndex    = 2
	limeIndex   = 3
	blueIndex   = 4
	yellowIndex = 5
)

var palette = []color.Color{
	whiteIndex:  color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	blackIndex:  color.RGBA{0x00, 0x00, 0x00, 0xFF},
	redIndex:    color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	limeIndex:   color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	blueIndex:   color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	yellowIndex: color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
}

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

		minColorIndex = 1
		maxColorIndex = 5
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	freq := r.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		randomIndex := rand.Intn(maxColorIndex) + minColorIndex

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randomIndex))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return gif.EncodeAll(out, &anim)
}
