package main

type Board [cols][rows]bool

func (board *Board) alive(i, j int) {
	board[i][j] = true
}

func (board *Board) kill(i, j int) {
	board[i][j] = false
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

func (board *Board) liveNeighborCount(i, j int) int {
	count := 0
	for _, dir := range dirs {
		if i <= 0 && dir[0] == -1 ||
			i >= cols-1 && dir[0] == 1 ||
			j <= 0 && dir[1] == -1 ||
			j >= rows-1 && dir[1] == 1 {
			continue
		}
		if board[i+dir[0]][j+dir[1]] {
			count++
		}
	}
	return count
}

func (board *Board) update() {
	var newBoard [cols][rows]bool
	for i := range board {
		for j := range board[0] {
			newBoard[i][j] = board[i][j]
			liveNeighbors := board.liveNeighborCount(i, j)
			if liveNeighbors < 2 || liveNeighbors > 3 {
				newBoard[i][j] = false
			} else if liveNeighbors == 3 {
				newBoard[i][j] = true
			}
		}
	}
	*board = newBoard
}
