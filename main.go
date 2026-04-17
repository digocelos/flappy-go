package main

import (
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 800
	screenHeight = 640
	birdWidth    = 38
	birdHeight   = 26
	gravityAccel = 0.3
	jumpImpulse  = -6.2
	pipeWidth    = 60
	pipeGap      = 150
	pipeSpeed    = 2.5
	pipeSpacing  = 220
	pipePadding  = 60
)

var (
	background = color.RGBA{R: 40, G: 120, B: 200, A: 0xff}
	birdColor  = color.RGBA{R: 250, G: 200, B: 30, A: 0xff}
	pipeColor  = color.RGBA{R: 20, G: 150, B: 40, A: 0xff}
)

const displayFrames = 60

var autoExit = os.Getenv("FLAPPY_AUTO_EXIT") == "1"

type Game struct {
	frames    int
	birdX     float64
	birdYPos  float64
	birdYVel  float64
	pipes     []Pipe
	nextPipeX float64
	gameOver  bool
}

func NewGame() *Game {
	g := &Game{
		birdX:    screenWidth/2 - birdWidth/2,
		birdYPos: screenHeight/2 - birdHeight/2,
	}
	g.initPipes()
	return g
}

type Pipe struct {
	x    float64
	gapY float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newPipe(x float64) Pipe {
	maxGapY := float64(screenHeight) - pipeGap - pipePadding
	if maxGapY < pipePadding {
		maxGapY = pipePadding
	}
	gapY := pipePadding + rand.Float64()*(maxGapY-pipePadding)
	return Pipe{x: x, gapY: gapY}
}

func (g *Game) initPipes() {
	g.pipes = g.pipes[:0]
	g.nextPipeX = float64(screenWidth) + pipeSpacing
	for i := 0; i < 3; i++ {
		g.pipes = append(g.pipes, newPipe(g.nextPipeX))
		g.nextPipeX += pipeSpacing
	}
}

func (g *Game) Update() error {
	if autoExit {
		g.frames++
		if g.frames > displayFrames {
			os.Exit(0)
		}
	}

	if g.gameOver {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.birdYVel = jumpImpulse
	}

	g.birdYVel += gravityAccel
	g.birdYPos += g.birdYVel

	for i := range g.pipes {
		g.pipes[i].x -= pipeSpeed
	}

	for len(g.pipes) > 0 && g.pipes[0].x+pipeWidth < 0 {
		g.pipes = g.pipes[1:]
	}

	for len(g.pipes) == 0 || g.pipes[len(g.pipes)-1].x+pipeWidth <= float64(screenWidth) {
		g.pipes = append(g.pipes, newPipe(g.nextPipeX))
		g.nextPipeX += pipeSpacing
	}

	if g.checkCollision() {
		g.gameOver = true
	}
	return nil
}

func (g *Game) checkCollision() bool {
	birdX1 := g.birdX
	birdY1 := g.birdYPos
	birdX2 := birdX1 + birdWidth
	birdY2 := birdY1 + birdHeight

	if birdY1 <= 0 || birdY2 >= float64(screenHeight) {
		return true
	}

	for _, pipe := range g.pipes {
		topX1 := pipe.x
		topY1 := float64(0)
		topX2 := topX1 + pipeWidth
		topY2 := pipe.gapY
		if overlap(birdX1, birdY1, birdX2, birdY2, topX1, topY1, topX2, topY2) {
			return true
		}

		bottomX1 := pipe.x
		bottomY1 := pipe.gapY + pipeGap
		bottomX2 := bottomX1 + pipeWidth
		bottomY2 := float64(screenHeight)
		if overlap(birdX1, birdY1, birdX2, birdY2, bottomX1, bottomY1, bottomX2, bottomY2) {
			return true
		}
	}

	return false
}

func overlap(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
	return ax1 < bx2 && ax2 > bx1 && ay1 < by2 && ay2 > by1
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background)
	ebitenutil.DrawRect(screen, g.birdX, g.birdYPos, birdWidth, birdHeight, birdColor)
	for _, pipe := range g.pipes {
		topHeight := pipe.gapY
		bottomY := pipe.gapY + pipeGap
		topHeightRect := topHeight
		bottomHeight := float64(screenHeight) - bottomY
		ebitenutil.DrawRect(screen, pipe.x, 0, pipeWidth, topHeightRect, pipeColor)
		ebitenutil.DrawRect(screen, pipe.x, bottomY, pipeWidth, bottomHeight, pipeColor)
	}
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
