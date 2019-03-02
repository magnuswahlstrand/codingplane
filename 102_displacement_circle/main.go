package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"golang.org/x/image/colornames"

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

func update(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.3)
	memory.DrawImage(darken, op)

	r := 0.8
	wt := 2 * math.Pi * t
	rx := r * math.Cos(wt)
	ry := r * math.Sin(wt)
	scale := 0.008 / (w / 600.0)
	length := w / 2.0
	for _, pixel := range pixels {
		dPos := pixel.pos

		xx := dPos.X + pixel.amplitude*length*noiser.Noise4D(scale*dPos.X, scale*dPos.Y, rx, ry) + 50
		yy := dPos.Y + pixel.amplitude*length*noiser.Noise4D(100+scale*dPos.X, scale*dPos.Y, rx, ry) + 50
		memory.Set(int(xx), int(yy), gfx.ColorWithAlpha(colornames.White, 255))
	}

	screen.DrawImage(memory, &ebiten.DrawImageOptions{})

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v", ebiten.CurrentTPS()))

	t += 0.5 / 60.0
	return nil
}

var (
	t               float64
	imgBytes        []byte
	motionBlurSteps = 3
	pixels          []Pixel
	images          = []*ebiten.Image{}
	darken          *ebiten.Image
	memory          *ebiten.Image
)

type Pixel struct {
	pos       gfx.Vec
	amplitude float64
}

func main() {

	tmp, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	tmp.Fill(color.Black)
	darken = tmp
	tmp, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	memory = tmp
	memory.Fill(color.Transparent)

	center := gfx.V(0.5, 0.5)
	for i := 0; i < 30000; i++ {
		pos := gfx.V(rand.Float64(), rand.Float64())
		d := pos.To(center).Len()
		if d < 0.5 {
			pixels = append(pixels, Pixel{
				pos:       pos.Scaled(w - 100),
				amplitude: math.Sin(math.Pi*(0.5-d)) / 2, // from 0.5 --> 0
			})
		}
	}

	ebiten.SetMaxTPS(30)

	if err := ebiten.Run(update, w, h, 0.7, "displacement circle"); err != nil {
		fmt.Println("exited")
	}
}
