package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	cellWidth    = 20
	windowHeight = 900
	windowWidth  = 1500
	rows         = windowHeight / cellWidth
	cols         = windowWidth / cellWidth
)

type Game struct {
	board Board
	pause bool
}

func (g *Game) Update() error {
	if g.pause {
	} else {
		g.board.update()
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for i := range cols {
		for j := range rows {
			if g.board[i][j] {
				vector.DrawFilledRect(
					screen,
					float32(i*cellWidth), float32(j*cellWidth), float32(cellWidth), float32(cellWidth),
					color.Black, true,
				)
			}
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{pause: true}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetWindowFloating(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

