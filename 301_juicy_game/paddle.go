package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
)

type Paddle struct {
	pos gfx.Vec
}

func NewPaddle() Paddle {
	offsetYBottom := 50.0
	x := width / 2.0
	y := height - offsetYBottom
	return Paddle{
		pos: gfx.V(x, y),
	}
}

func (p *Paddle) draw(screen *ebiten.Image) {
	offsetX := float64(paddleSize.Dx()) / 2
	ebitenutil.DrawRect(screen, p.pos.X-offsetX, p.pos.Y, float64(paddleSize.Dx()), float64(paddleSize.Dy()), paddleColor)
}

func (p *Paddle) updatePosition(x float64) {
	p.pos.X = x
}
