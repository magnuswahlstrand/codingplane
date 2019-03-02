package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
)

const (
	width        = 400
	height       = 400
	circleRadius = 10
)

type Circle struct {
	gfx.Vec
	scale float64
}

var circles = []Circle{}
var circleImg *ebiten.Image
var noiser = gfx.NewSimplexNoise(0)

func update(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	for i, c := range circles {
		t := float64(time.Now().UnixNano()/1000%10000000) / 10000000
		sx := 0.5 + 0.5*noiser.Noise3D(math.Cos(2*math.Pi*t), math.Sin(2*math.Pi*t), float64(i)/float64(len(circles)))
		sy := 0.5 + 0.5*noiser.Noise3D(math.Cos(2*math.Pi*t), math.Sin(2*math.Pi*t), float64(i+len(circles))/float64(len(circles)))
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(c.scale, c.scale)
		op.GeoM.Translate(width*sx, height*sy)
		op.ColorM.Scale(1, 1, 1, 0.8)
		screen.DrawImage(circleImg, op)
	}
	return nil
}

func main() {
	tmpImg := gfx.NewImage(2*circleRadius, 2*circleRadius, color.Transparent)
	gfx.DrawCircle(tmpImg, gfx.V(circleRadius, circleRadius), circleRadius, 0, colornames.Pink)
	circleImg, _ = ebiten.NewImageFromImage(tmpImg, ebiten.FilterDefault)

	for i := 0; i < 10; i++ {
		pos := gfx.V(width*rand.Float64()-2*circleRadius, height*rand.Float64()-2*circleRadius)

		circles = append(circles, Circle{Vec: pos, scale: (0.5 + 0.5*rand.Float64())})
	}

	if err := ebiten.Run(update, width, height, 0.7, "pink arm"); err != nil {
		log.Fatal(err)
	}
}
