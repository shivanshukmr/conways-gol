package main

type World [cols][rows]bool

func (world *World) alive(i, j int) {
	world[i][j] = true
}

func (world *World) kill(i, j int) {
	world[i][j] = false
}

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

func (world *World) liveNeighborCount(i, j int) int {
	count := 0
	for _, dir := range dirs {
		if i <= 0 && dir[0] == -1 ||
			i >= cols-1 && dir[0] == 1 ||
			j <= 0 && dir[1] == -1 ||
			j >= rows-1 && dir[1] == 1 {
			continue
		}
		if world[i+dir[0]][j+dir[1]] {
			count++
		}
	}
	return count
}

func (world *World) update() {
	var newWorld [cols][rows]bool
	for i := range world {
		for j := range world[0] {
			newWorld[i][j] = world[i][j]
			liveNeighbors := world.liveNeighborCount(i, j)
			if liveNeighbors < 2 || liveNeighbors > 3 {
				newWorld[i][j] = false
			} else if liveNeighbors == 3 {
				newWorld[i][j] = true
			}
		}
	}
	*world = newWorld
}
