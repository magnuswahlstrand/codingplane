package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

const (
	width, height   = 400, 400
	radius          = 10
	fade            = 0.92
	motionBlurSteps = 40
)

var images = []*ebiten.Image{}

func update(screen *ebiten.Image) error {
	images[0].Clear()
	images = append(images[1:], images[0])

	t := float64(time.Now().UnixNano() / int64(time.Millisecond))
	amplitude := width/2.0 - radius - 40

	// Add noise
	wt := 2 * math.Pi * t / 4000
	x := math.Cos(wt)
	y := math.Sin(wt)
	amplitude += 20 * noiser.Noise2D(x, 1.1*y)
	// noiseX := noiser.Noise2D(math.Pi, 1.1*y)
	// noiseY := noiser.Noise2D(math.Pi, 1.1*y)

	op := &ebiten.DrawImageOptions{}
	sx := amplitude*x + width/2 - radius
	sy := amplitude*y + height/2 - radius
	op.GeoM.Translate(sx, sy)
	images[len(images)-1].DrawImage(circleImg, op)

	for i, img := range images {
		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(1, 1, 1, math.Pow(fade, float64(len(images)-i)))
		screen.DrawImage(img, op)
	}

	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Current TPS: %v", ebiten.CurrentTPS()))
	return nil
}

var circleImg *ebiten.Image
var noiser = gfx.NewSimplexNoise(0)

func main() {

	tmp := gfx.NewImage(2*radius, 2*radius, color.Transparent)
	gfx.DrawCircle(tmp, gfx.V(radius, radius), radius, 0, colornames.Red)
	circleImg, _ = ebiten.NewImageFromImage(tmp, ebiten.FilterDefault)

	var sum float64
	for i := 0; i < motionBlurSteps; i++ {
		tmp, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
		images = append(images, tmp)
		sum += math.Pow(fade, float64(i))
	}

	if err := ebiten.Run(update, width, height, 0.7, "motion blur"); err != nil {
		log.Fatal(err)
	}
}
