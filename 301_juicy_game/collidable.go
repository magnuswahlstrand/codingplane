package main

import (
	"github.com/peterhellberg/gfx"
)

type Collidable interface {
	Hitbox() gfx.Rect
	MarkCollided(bool)
}
