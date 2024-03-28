package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	windowHeight = 900
	windowWidth  = 1500
)

const cellWidth = 20

const (
	rows = windowHeight / cellWidth
	cols = windowWidth / cellWidth
)

var board = [cols][rows]bool{}
var newBoard = [cols][rows]bool{}

var dirs = [8][2]int{
	{-1, -1}, // top left
	{-1, 0},  // left
	{-1, 1},  // bottom left
	{0, -1},  // top middle
	{0, 1},   // bottom middle
	{1, -1},  // top right
	{1, 0},   // right
	{1, 1},   // bottom right
}

func a(i, j int) {
	board[i][j] = true
}

func getLiveNeighborCount(i, j int) int {
	count := 0

	for _, dir := range dirs {
		if i <= 0 && dir[0] == -1 {
			continue
		}
		if i >= cols - 1 && dir[0] == 1 {
			continue
		}
		if j <= 0 && dir[1] == -1 {
			continue
		}
		if j >= rows - 1 && dir[1] == 1 {
			continue
		}
		if board[i + dir[0]][j + dir[1]] {
			count++
		}
	}
	return count
}

type Game struct{}

func (g *Game) Update() error {
	time.Sleep(50 * time.Millisecond)
	for i := range board {
		for j := range board[0] {
			newBoard[i][j] = board[i][j]
			liveNeighbors := getLiveNeighborCount(i, j)
			if liveNeighbors < 2 || liveNeighbors > 3 {
				newBoard[i][j] = false
			} else if liveNeighbors == 3 {
				newBoard[i][j] = true
			}
		}
	}
	board = newBoard
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := range cols {
		for j := range rows {
			if board[i][j] {
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
	game := &Game{}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetWindowFloating(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

