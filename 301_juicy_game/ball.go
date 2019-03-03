package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
)

const (
	startingVelocity = 7
)

type Ball struct {
	pos      gfx.Vec
	velocity gfx.Vec
}

func (b *Ball) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.pos.X-float64(ballSize.Dx())/2, b.pos.Y-float64(ballSize.Dy())/2, float64(ballSize.Dx()), float64(ballSize.Dy()), ballColor)
}

func newBall() Ball {
	offsetYBottom := 130.0
	x := width / 2.0
	y := height - offsetYBottom
	fmt.Println(x, y)
	return Ball{
		pos:      gfx.V(x, y),
		velocity: gfx.V(1, 0).Rotated(2 * math.Pi * rand.Float64()).Scaled(startingVelocity),
	}
}

func (b *Ball) updatePosition() {
	b.pos = b.pos.Add(b.velocity)
}
