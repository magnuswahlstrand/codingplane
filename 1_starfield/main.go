package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
)

type Star struct {
	pos gfx.Vec
	z   float64
}

const (
	width  = 400
	height = 400
)

var board = gfx.R(5, 5, width-5, height-5)
var offset = gfx.V(width/2, height/2)

func update(screen *ebiten.Image) error {

	for i, s := range stars {

		stars[i].z -= 2
		v := s.pos
		v = v.Sub(offset)
		v = v.Scaled(width / s.z)
		v = v.Add(offset)

		if board.Contains(v) {
			gfx.DrawCicleFast(screen, v, 1+2*(width-s.z)/width, gfx.ColorWithAlpha(color.White, uint8(100+155*(1-s.z/width))))
		} else {
			// Replace star
			stars[i] = newStar()
		}
	}

	return nil
}

var stars []Star

func newStar() Star {
	return Star{
		gfx.IV(rand.Intn(width), rand.Intn(width)),
		width,
	}
}

func main() {
	for i := 0; i < 100; i++ {
		stars = append(stars, newStar())
	}

	if err := ebiten.Run(update, width, height, 1, "menu example"); err != nil {
		log.Fatal(err)
	}
}
