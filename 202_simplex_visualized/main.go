package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/colornames"

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

var white bool

func update(screen *ebiten.Image) error {
	scale := 0.02

	for i := 0; i < 3*w; i++ {
		baseline := 100
		x := i / 3
		y := baseline + int(50*noiser.Noise2D(scale*float64(i)/3, t))

		// Baseline
		screen.Set(x, baseline, colornames.Gray)

		// Noise line
		screen.Set(x, y, color.White)
	}

	y := 200
	white = false
	var start, finish gfx.Vec
	for i := 0; i < w; i++ {
		r := noiser.Noise2D(scale*float64(i), t)

		switch {
		case r > 0 && !white:
			start = gfx.IV(i, y)
			white = true

		case r < 0 && white:
			finish = gfx.IV(i, y+30)
			white = false

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(start.X, start.Y)
			screen.DrawImage(img.SubImage(gfx.IR(int(start.X), int(start.Y), int(finish.X), int(finish.Y))).(*ebiten.Image), op)
		}
	}

	t += 0.3 / 20.0
	return nil
}

func main() {
	tmp, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	tmp.Fill(color.White)
	img = tmp

	noiser = gfx.NewSimplexNoise(0)
	// for i := 0; i < 50; i++ {

	// 	r := gfx.IR(0, 0, rand.Intn(100), rand.Intn(100))
	// 	rectangles = append(rectangles, Rectangle{
	// 		pos:  gfx.IV(rand.Intn(w-r.Dx()-100), rand.Intn(h-r.Dy()-100)),
	// 		size: r,
	// 	})
	// }

	if err := ebiten.Run(update, w, h, 1, "kinect example"); err != nil {
		fmt.Println("exited")
	}
}
