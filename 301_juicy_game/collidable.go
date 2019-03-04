package main

import (
	"github.com/peterhellberg/gfx"
)

type Collidable interface {
	CollidingWith(gfx.Vec) bool
}
