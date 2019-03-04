package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

type Walls struct {
	top, left, right, bottom gfx.Rect
	collided                 bool
}

var borderWidth = 10.0

func newWalls() Walls {
	top := gfx.R(0, 0, width, borderWidth)
	left := gfx.R(0, borderWidth, borderWidth, height)
	right := gfx.R(width-borderWidth, borderWidth, width, height)
	bottom := gfx.R(0, height, width, height+borderWidth)

	return Walls{
		top:    top,
		left:   left,
		right:  right,
		bottom: bottom,
	}
}

func (w *Walls) CollidingWith(pos gfx.Vec) bool {
	w.collided = false
	for _, hitbox := range []gfx.Rect{w.top, w.left, w.right, w.bottom} {
		if hitbox.Contains(pos) {
			w.collided = true
			break
		}
	}
	return w.collided
}

func (w *Walls) draw(screen *ebiten.Image) {
	screen.DrawImage(colorImg.SubImage(w.top.Bounds()).(*ebiten.Image), &ebiten.DrawImageOptions{})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(borderWidth))
	screen.DrawImage(colorImg.SubImage(w.left.Bounds()).(*ebiten.Image), op)
	op.GeoM.Translate(w.right.Min.X, 0)
	screen.DrawImage(colorImg.SubImage(w.right.Bounds()).(*ebiten.Image), op)
}
