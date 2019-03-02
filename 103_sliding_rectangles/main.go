package main

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

const (
	w, h    = 440, 440
	padding = 40
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

	var x int
	for j, col := range columns {

		wt := 2 * math.Pi * t
		px := math.Cos(wt + float64(j*10))
		py := math.Sin(wt)

		white = false
		var start, finish gfx.Vec
		for i := 0; i < h-2*padding; i++ {
			r := noiser.Noise3D(col.scale*float64(i), col.radius*px, col.radius*py)

			switch {
			case r > 0 && !white:
				start = gfx.IV(x, i)
				white = true

			case r < 0 && white:
				finish = gfx.IV(x+col.width, i)
				white = false

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(start.X+padding, start.Y+padding)
				screen.DrawImage(img.SubImage(gfx.IR(int(start.X), int(start.Y), int(finish.X), int(finish.Y))).(*ebiten.Image), op)
			}
		}
		x += col.width
	}

	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeXor
	screen.DrawImage(circleImg, op)

	size := w - 2.0*padding
	ebitenutil.DrawLine(screen, padding, padding, padding, padding+size, colornames.Silver)
	ebitenutil.DrawLine(screen, padding, padding, padding+size, padding, colornames.Silver)
	ebitenutil.DrawLine(screen, padding, padding+size, padding+size, padding+size, colornames.Silver)
	ebitenutil.DrawLine(screen, padding+size, padding, padding+size, padding+size, colornames.Silver)
	t += 0.10 / 60.0
	return nil
}

var (
	t          float64
	noiser     *gfx.SimplexNoise
	img        *ebiten.Image
	circleImg  *ebiten.Image
	rectangles []Rectangle
	columns    []Column
)

type Column struct {
	width  int
	scale  float64
	radius float64
}

func main() {
	shapeColor := colornames.Silver
	tmp, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	tmp.Fill(shapeColor)
	img = tmp

	tmp2, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	gfx.DrawCircle(tmp2, gfx.V(w/2, h/2), (w-80)/2, 0, shapeColor)
	circleImg = tmp2

	columns = []Column{
		Column{width: 10, scale: 0.02, radius: 0.3},
		Column{width: 30, scale: 0.5, radius: 0.01},
		Column{width: 20, scale: 0.01, radius: 0.5},
		Column{width: 60, scale: 0.04, radius: 1},
		Column{width: 5, scale: 0.5, radius: 0.02},
		Column{width: 35, scale: 0.15, radius: 0.65},
		Column{width: 25, scale: 0.02, radius: 0.28},
		Column{width: 3, scale: 0.01, radius: 1.0},
		Column{width: 65, scale: 0.05, radius: 1},
		Column{width: 15, scale: 0.5, radius: 0.1},
		Column{width: 10, scale: 0.02, radius: 2},
		Column{width: 30, scale: 0.01, radius: 2},
		Column{width: 20, scale: 0.05, radius: 0.2},
		Column{width: 32, scale: 0.1, radius: 0.5},
	}

	noiser = gfx.NewSimplexNoise(0)

	if err := ebiten.Run(update, w, h, 1, "kinect example"); err != nil {
		fmt.Println("exited")
	}
}
