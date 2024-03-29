package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.pause = !g.pause
	}

	if g.pause {
		leftMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		rightMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

		if leftMousePressed || rightMousePressed {
			cursorx, cursory := ebiten.CursorPosition()
			if cursorx < 0 || cursorx > windowWidth-1 || cursory < 0 || cursory > windowHeight-1 {
				return nil
			}
			if leftMousePressed {
				g.board.alive(
					cursorx/cellWidth,
					cursory/cellWidth,
				)
			}
			if rightMousePressed {
				g.board.kill(
					cursorx/cellWidth,
					cursory/cellWidth,
				)
			}
		}
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
