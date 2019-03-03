package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

type Brick struct {
	pos      gfx.Vec
	collided bool
}

func (b *Brick) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.pos.X, b.pos.Y)

	if b.collided {
		op.ColorM.Scale(100, 100, 100, 1)
	}
	screen.DrawImage(colorImg.SubImage(brickSize.Bounds()).(*ebiten.Image), op)

}

func (b *Brick) Hitbox() gfx.Rect {
	return brickSize.Moved(b.pos) //.Moved(paddleSize.Max.Scaled(-0.5))
}

func (b *Brick) MarkCollided(t bool) {
	b.collided = t
}
