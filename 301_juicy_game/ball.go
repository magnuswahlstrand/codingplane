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
	startingVelocity = 1
)

type Ball struct {
	pos      gfx.Vec
	velocity gfx.Vec
}

func (b *Ball) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.pos.X-ballSize.W()/2, b.pos.Y-ballSize.W()/2, ballSize.W(), ballSize.W(), ballColor)
}

func newBall() Ball {
	offsetYBottom := 130.0
	x := width / 2.0
	y := height - offsetYBottom
	return Ball{
		pos: gfx.V(x, y),
		// velocity: gfx.V(1, 0).Rotated(2 * math.Pi * rand.Float64()).Scaled(startingVelocity),
		velocity: gfx.V(0, 1).Rotated(0 * 2 * math.Pi * rand.Float64()).Scaled(startingVelocity),
	}
}

func (b *Ball) updatePosition() {
	// Move y
	b.pos.Y += b.velocity.Y

	// Check for collision
	for _, c := range collidableObjects {

		fmt.Println(c.Hitbox())
		c.MarkCollided(c.Hitbox().Contains(b.pos))
	}

	// Check for collision

	b.pos = b.pos.Add(b.velocity)
	// Move x
}
