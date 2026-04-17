package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 480
	screenHeight = 640
	birdWidth    = 38
	birdHeight   = 26
	gravityAccel = 0.3
)

var background = color.RGBA{R: 40, G: 120, B: 200, A: 0xff}
var birdColor = color.RGBA{R: 250, G: 200, B: 30, A: 0xff}

const displayFrames = 60

var autoExit = os.Getenv("FLAPPY_AUTO_EXIT") == "1"

type Game struct {
	frames   int
	birdX    float64
	birdYPos float64
	birdYVel float64
}

func NewGame() *Game {
	return &Game{
		birdX:    screenWidth/2 - birdWidth/2,
		birdYPos: screenHeight/2 - birdHeight/2,
	}
}

func (g *Game) Update() error {
	if autoExit {
		g.frames++
		if g.frames > displayFrames {
			os.Exit(0)
		}
	}

	g.birdYVel += gravityAccel
	g.birdYPos += g.birdYVel
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	ebitenutil.DrawRect(screen, g.birdX, g.birdYPos, birdWidth, birdHeight, birdColor)
}

func (Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Flappy Go")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
