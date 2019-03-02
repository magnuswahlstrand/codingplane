package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"golang.org/x/image/colornames"

	"github.com/fogleman/ease"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func limit(op *ebiten.DrawImageOptions, r, g, b, limit float64) {
	op.ColorM.Scale(min(max(r, limit), 1), min(max(g, limit), 1), min(max(b, limit), 1), 1)
}

func colorFunc(op *ebiten.DrawImageOptions, t2 float64) {
	switch {
	case t2 < 1:
		// White --> Blue
		limit(op, 1-2*t2, 1-2*t2, 1, 0.7)
	case t2 < 2:
		limit(op, 0, 0, 1, 0.7)
	case t2 < 3:
		// Blue --> Green
		ot := t2 - 2
		limit(op, 0, 2*ot, 2*(1-2*ot), 0.7)
	case t2 < 4:
		// Green
		limit(op, 0, 1, 0, 0.7)
	case t2 < 5:
		// Green --> Blue
		ot := t2 - 4
		limit(op, 2*ot, 2*(1-2*ot), 0, 0.7)
	case t2 < 6:
		// Red
		limit(op, 1, 0, 0, 0.7)
	case t2 < 7:
		// Red --> White
		ot := t2 - 6
		limit(op, 1, 2*ot, 2*ot, 0.7)
	}
}

func redBlue(op *ebiten.DrawImageOptions, t2 float64) {
	switch {
	case t2 < 1:
		limit(op, 1-2*t2, 1-2*t2, 1, 0.7)
	case t2 < 2:
		limit(op, 0, 0, 1, 0.7)
	case t2 < 3:
		limit(op, 1, 1-2*t2, 1-2*t2, 0.7)
	case t2 < 4:
		limit(op, 1, 0, 0, 0.7)
	}
}

func blackWhite(op *ebiten.DrawImageOptions, t2 float64) {
	switch {
	case t2 < 1:
		limit(op, 1-t2, 1-t2, 1-t2, 0.7)
	case t2 < 2:
		limit(op, 0, 0, 0, 0.7)
	case t2 < 3:
		ot := t2 - 2
		limit(op, 0.7-ot, 0.7-ot, 0.7-ot, 0.5)
	case t2 < 4:
		limit(op, 0, 0, 0, 0.5)
	case t2 < 5:
		ot := t2 - 4
		limit(op, 0.5-ot, 0.5-ot, 0.5-ot, 0.3)
	case t2 < 6:
		limit(op, 0, 0, 0, 0.3)
	case t2 < 7:
		ot := t2 - 6
		limit(op, 0.7+ot, 0.7+ot, 0.7+ot, 0.3)
	case t2 < 8:
		limit(op, 1, 1, 1, 0.5)

	}
}

func update(screen *ebiten.Image) error {

	// Draw grid
	step := 100.0
	for lx := 0.0 - step*t2; lx < w; lx += step {
		ebitenutil.DrawLine(screen, math.Floor(lx), 0, math.Floor(lx), h, gfx.ColorWithAlpha(colornames.Blanchedalmond, 30))
	}

	for i := 0; i < nBricks; i++ {
		width := w - float64(brick.Dx())

		var x float64
		if t < 0.5 {
			x = ease.InOutQuint(2*0.8*t+brickOffset[i]) - t/2
		} else {
			x = 1 - t/2
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(padding+width*x, float64(i*(brick.Dy()+1)))

		blackWhite(op, t2)

		screen.DrawImage(colorImage.SubImage(brick).(*ebiten.Image), op)
	}

	dt := 0.5 / 60
	t += dt
	t2 += dt
	if t > 2 {
		t = 0
		brickOffset = newOffsets()
	}

	if t2 > 8 {
		t2 = 0
	}

	return nil

}

const (
	w, h    = 500, 500
	padding = 50.0
	nBricks = 50
)

var (
	brick       image.Rectangle
	colorImage  *ebiten.Image
	t, t2       float64
	brickOffset []float64
)

func newOffsets() []float64 {
	var offsets []float64
	for i := 0; i < nBricks; i++ {
		offsets = append(offsets, 0.2*rand.Float64())
	}
	return offsets
}

func main() {
	colorImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	colorImage.Fill(color.White)

	brick = gfx.IR(0, 0, 30, 10)

	brickOffset = newOffsets()

	if err := ebiten.Run(update, w, h, 1, "moving bricks"); err != nil {
		fmt.Println("exited")
	}
}
