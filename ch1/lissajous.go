package ch1

import (
	"image/color"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

var palette = []color.Color{color.White, color.Black, red, green, blue}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	redIndex = 2
	greenIndex = 3
	blueIndex = 4
)

func Lissajous(out io.Writer, cycles int, color string) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	colorIndex := uint8(whiteIndex)
	switch color {
	case "red":
		colorIndex = redIndex
	case "green":
		colorIndex = greenIndex
	case "blue":
		colorIndex = blueIndex
	default:
		colorIndex = blackIndex
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
