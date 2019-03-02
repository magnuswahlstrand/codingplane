package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

// https://necessarydisorder.wordpress.com/2017/11/15/drawing-from-noise-and-then-making-animated-loopy-gifs-from-there/

const (
	w, h = 400, 400
)

var (
	img       *image.RGBA
	ebitenImg *ebiten.Image
	noiser    *gfx.SimplexNoise
)

func init() {
	noiser = gfx.NewSimplexNoise(0)
	img = gfx.NewImage(w, h, color.Transparent)
	tmp, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	ebitenImg = tmp
}

func mp(v, inMin, inMax, outMin, outMax float64) byte {
	rangeIn := inMax - inMin
	rangeOut := outMax - outMin

	return byte(rangeOut*(v-inMin)/rangeIn + outMin)
}

func Sigmoid(v float64) float64 {
	return 1.0 / (1.0 + math.Exp(-v))
}

func update(screen *ebiten.Image) error {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			radius := 1.0
			scale := 0.02
			// t := float64(time.Now().UnixNano() / int64(time.Millisecond))
			wt := 2 * math.Pi * t
			rx := math.Cos(wt)
			ry := math.Sin(wt) //float64(time.Now().UnixNano() / int64(time.Millisecond)))
			v := noiser.Noise4D(scale*float64(x), scale*float64(y), radius*rx, radius*ry)
			i := 4 * (y*w + x)
			if v > 0 {
				imgBytes[i] = 255
				imgBytes[i+1] = 255
				imgBytes[i+2] = 255

			} else {
				imgBytes[i] = 0
				imgBytes[i+1] = 0
				imgBytes[i+2] = 0
			}
		}
	}
	ebitenImg.ReplacePixels(imgBytes)
	screen.DrawImage(ebitenImg, &ebiten.DrawImageOptions{})
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v", ebiten.CurrentTPS()))

	t += 0.1 / 40.0
	return nil
}

var (
	t        float64
	imgBytes []byte
)

func main() {
	imgBytes = make([]byte, 4*w*h)
	for i := 0; i < w*h; i++ {
		imgBytes[4*i] = 255
		imgBytes[4*i+3] = 255
	}

	ebiten.SetMaxTPS(40)

	if err := ebiten.Run(update, w, h, 0.7, "kinect example"); err != nil {
		fmt.Println("exited")
	}
}
