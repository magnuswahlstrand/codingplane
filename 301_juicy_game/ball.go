package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/peterhellberg/gfx"
)

const (
	startingVelocity = 5
)

type Ball struct {
	pos      gfx.Vec
	velocity gfx.Vec
}

func (b *Ball) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.pos.X-ballSize.W()/2, b.pos.Y-ballSize.W()/2, ballSize.W(), ballSize.H(), ballColor)
}

func newBall() Ball {
	offsetYBottom := 130.0
	x := width / 2.0
	y := height - offsetYBottom
	return Ball{
		pos:      gfx.V(x, y),
		velocity: gfx.V(1, 0).Rotated(2 * math.Pi * rand.Float64()).Scaled(startingVelocity),
		// velocity: gfx.V(0, 1).Rotated(0 * 2 * math.Pi * rand.Float64()).Scaled(startingVelocity),
	}
}

func (b *Ball) updatePosition() {
	var collidedX, collidedY bool

	// Move y
	b.pos.Y += b.velocity.Y

	// Check for collision
	for _, c := range collidableObjects {
		collided := c.CollidingWith(b.pos)
		collidedY = collidedY || collided
	}

	if collidedY {
		// Revert move, and change direction
		b.pos.Y -= b.velocity.Y
		b.velocity.Y = -b.velocity.Y
	}

	// Move x
	b.pos.X += b.velocity.X

	// Check for collision
	for _, c := range collidableObjects {
		collided := c.CollidingWith(b.pos)
		collidedX = collidedX || collided
	}

	if collidedX {
		// Revert move, and change direction
		b.pos.X -= b.velocity.X
		b.velocity.X = -b.velocity.X
	}
}
