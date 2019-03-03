package main

import (
	"image"
	"log"

	"github.com/peterhellberg/gfx"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
)

const (
	width   = 600
	height  = 400
	bricksX = 10
	bricksY = 8
	spacing = 8
)

var (
	brickSize  gfx.Rect = gfx.R(0, 0, 35, 15)
	paddleSize gfx.Rect = gfx.R(0, 0, 70, 25)
	ballSize   gfx.Rect = gfx.R(0, 0, 15, 15)
)

func drawBorder(screen *ebiten.Image) {
	borderWidth := 10
	top := image.Rect(0, 0, width, borderWidth)
	side := image.Rect(0, borderWidth, borderWidth, height)
	screen.DrawImage(colorImg.SubImage(top).(*ebiten.Image), &ebiten.DrawImageOptions{})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(borderWidth))
	screen.DrawImage(colorImg.SubImage(side).(*ebiten.Image), op)
	op.GeoM.Translate(width-float64(borderWidth), 0)
	screen.DrawImage(colorImg.SubImage(side).(*ebiten.Image), op)
}

func drawBricks(screen *ebiten.Image) {
	offsetY := 50.0
	offsetX := (width - (brickSize.W()+spacing)*bricksX) / 2

	for y := 0.0; y < bricksY; y++ {
		for x := 0.0; x < bricksX; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate((brickSize.W()+spacing)*x+offsetX, (brickSize.H()+spacing)*y+offsetY)
			screen.DrawImage(colorImg.SubImage(brickSize.Bounds()).(*ebiten.Image), op)
		}
	}
}

func update(screen *ebiten.Image) error {

	// Update position
	x, _ := ebiten.CursorPosition()
	paddle.updatePosition(float64(x))
	ball.updatePosition()

	// op := &ebiten.DrawImageOptions{}
	screen.Fill(backgroundColor)
	drawBorder(screen)
	drawBricks(screen)
	ball.draw(screen)
	paddle.draw(screen)
	return nil
}

var (
	backgroundColor   = colornames.Lemonchiffon
	borderColor       = colornames.Darkseagreen
	paddleColor       = colornames.Firebrick
	ballColor         = colornames.Tomato
	colorImg          *ebiten.Image
	paddle            Paddle
	ball              Ball
	collidableObjects []Collidable
)

func main() {
	colorImg, _ = ebiten.NewImage(width+1, height+1, ebiten.FilterDefault)
	colorImg.Fill(borderColor)

	paddle = newPaddle()
	ball = newBall()

	collidableObjects = append(collidableObjects, &paddle)

	if err := ebiten.Run(update, width, height, 1, "juicy colors"); err != nil {
		log.Fatal(err)
	}
}
