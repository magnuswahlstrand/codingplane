package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

const (
	w, h = 400, 400
)

var (
	t          float64
	noiser     *gfx.SimplexNoise
	img        *ebiten.Image
	rectangles []Rectangle
)

type Pixel struct {
	pos       gfx.Vec
	amplitude float64
}

type Rectangle struct {
	pos  gfx.Vec
	size image.Rectangle
}

func update(screen *ebiten.Image) error {
	scale := 0.02
	radius := 0.2
	wt := 2 * math.Pi * t
	rx := radius * math.Cos(wt)
	ry := radius * math.Sin(wt)
	for _, r := range rectangles {
		op := &ebiten.DrawImageOptions{}

		xx := r.pos.X + 100*noiser.Noise4D(scale*r.pos.X, scale*r.pos.Y, rx, ry) + 50
		yy := r.pos.Y + 100*noiser.Noise4D(100+scale*r.pos.X, scale*r.pos.Y, rx, ry) + 50

		op.GeoM.Translate(xx, yy)
		op.CompositeMode = ebiten.CompositeModeXor
		screen.DrawImage(img.SubImage(r.size).(*ebiten.Image), op)
	}

	t += 0.3 / 60.0
	return nil
}

func main() {
	tmp, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	tmp.Fill(color.White)
	img = tmp

	noiser = gfx.NewSimplexNoise(0)
	for i := 0; i < 50; i++ {

		r := gfx.IR(0, 0, rand.Intn(100), rand.Intn(100))
		rectangles = append(rectangles, Rectangle{
			pos:  gfx.IV(rand.Intn(w-r.Dx()-100), rand.Intn(h-r.Dy()-100)),
			size: r,
		})
	}

	if err := ebiten.Run(update, w, h, 1, "kinect example"); err != nil {
		fmt.Println("exited")
	}
}
