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

func (b *Brick) CollidingWith(pos gfx.Vec) bool {
	hitbox := brickSize.Moved(b.pos)
	b.collided = hitbox.Contains(pos)
	return b.collided //.Moved(paddleSize.Max.Scaled(-0.5))
}
