package main

import (
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

func update(screen *ebiten.Image) error {

	// Update position
	x, _ := ebiten.CursorPosition()
	paddle.updatePosition(float64(x))
	ball.updatePosition()

	// op := &ebiten.DrawImageOptions{}
	screen.Fill(backgroundColor)
	walls.draw(screen)
	for _, b := range bricks {
		b.draw(screen)
	}
	paddle.draw(screen)
	ball.draw(screen)
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
	walls             Walls
	bricks            []*Brick
	collidableObjects []Collidable
)

func main() {
	colorImg, _ = ebiten.NewImage(width+1, height+1, ebiten.FilterDefault)
	colorImg.Fill(borderColor)

	paddle = newPaddle()
	collidableObjects = append(collidableObjects, &paddle)

	offsetY := 50.0
	offsetX := (width - (brickSize.W()+spacing)*bricksX) / 2
	for y := 0.0; y < bricksY; y++ {
		for x := 0.0; x < bricksX; x++ {
			sx := (brickSize.W()+spacing)*x + offsetX
			sy := (brickSize.H()+spacing)*y + offsetY

			b := Brick{pos: gfx.V(sx, sy)}
			bricks = append(bricks, &b)
			collidableObjects = append(collidableObjects, &b)
		}
	}

	walls = newWalls()
	collidableObjects = append(collidableObjects, &walls)

	ball = newBall()
	if err := ebiten.Run(update, width, height, 1, "juicy colors"); err != nil {
		log.Fatal(err)
	}
}
