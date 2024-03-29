package main

import "math"

func parseRle(rle string) World {
	var world World
	startPerc := 0.2
	startx, starty := int(math.Round(cols*startPerc)), int(math.Round(rows*startPerc))
	x, y := startx, starty

	repeat := 0
	for _, c := range rle {
		if c >= '0' && c <= '9' {
			repeat = repeat*10 + int(c-'0')
		} else if c == 'o' {
			world.alive(x, y)
			x++
			for i := 0; i < repeat-1; i++ {
				world.alive(x, y)
				x++
			}
			repeat = 0
		} else if c == 'b' {
			// world.kill(x, y)
			x++
			for i := 0; i < repeat-1; i++ {
				// world.kill(x, y)
				x++
			}
			repeat = 0
		} else if c == '$' {
			x = startx
			y++
			for i := 0; i < repeat-1; i++ {
				y++
			}
			repeat = 0
		}
	}

	return world
}
