package main

import (
	"bufio"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	fps          = 100
	cellWidth    = 2
	windowHeight = 900
	windowWidth  = 1500
	rows         = windowHeight / cellWidth
	cols         = windowWidth / cellWidth
)

type Game struct {
	world World
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
				g.world.alive(
					cursorx/cellWidth,
					cursory/cellWidth,
				)
			}
			if rightMousePressed {
				g.world.kill(
					cursorx/cellWidth,
					cursory/cellWidth,
				)
			}
		}
	} else {
		g.world.update()
		time.Sleep((1000 / fps) * time.Millisecond)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for i := range cols {
		for j := range rows {
			if g.world[i][j] {
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

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		var line string
		sc := bufio.NewScanner(file)

		for sc.Scan() {
			line = sc.Text()
			if line[0] != '#' && line[0] != ' ' && line[0] != '\t' {
				break
			}
		}
		// check pattern with boundaries
		// log.Println(sc.Text())

		var rle []string
		for sc.Scan() {
			rle = append(rle, sc.Text())
		}

		game.world = parseRle(strings.Join(rle, ""))

		file.Close()
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
