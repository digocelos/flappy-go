package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 480
	screenHeight = 640
)

var background = color.RGBA{R: 40, G: 120, B: 200, A: 0xff}

const displayFrames = 60

var autoExit = os.Getenv("FLAPPY_AUTO_EXIT") == "1"

type Game struct {
	frames int
}

func (g *Game) Update() error {
	if autoExit {
		g.frames++
		if g.frames > displayFrames {
			os.Exit(0)
		}
	}
	return nil
}

func (Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
}

func (Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Flappy Go")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
