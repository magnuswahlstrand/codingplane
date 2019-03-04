package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

type Paddle struct {
	pos      gfx.Vec
	collided bool
}

func newPaddle() Paddle {
	offsetYBottom := 50.0
	x := width / 2.0
	y := height - offsetYBottom
	return Paddle{
		pos: gfx.V(x, y),
	}
}

func (p *Paddle) draw(screen *ebiten.Image) {
	offsetX := paddleSize.W() / 2
	ebitenutil.DrawRect(screen, p.pos.X-offsetX, p.pos.Y, paddleSize.W(), paddleSize.H(), paddleColor)
	if p.collided {
		ebitenutil.DrawRect(screen, p.pos.X-offsetX, p.pos.Y, paddleSize.W(), paddleSize.H(), colornames.White)
	}
}

func (p *Paddle) updatePosition(x float64) {
	p.pos.X = x
}

func (p *Paddle) CollidingWith(pos gfx.Vec) bool {
	hitbox := paddleSize.Moved(p.pos).Moved(paddleSize.Max.ScaledXY(gfx.V(-0.5, 0)))
	p.collided = hitbox.Contains(pos)
	return p.collided //.Moved(paddleSize.Max.Scaled(-0.5))
}
