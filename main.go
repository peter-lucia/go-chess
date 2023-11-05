package main

import (
	"fmt"
)

type Cell int

const (
	Empty Cell = iota
	Queen
	King
	Bishop
	Horse
	Rook
	Pawn
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Board struct {
	state [8][8]Cell
}

type Piece struct {
	row         int
	col         int
	moveUpDx    int
	moveUpDy    int
	moveLeftDx  int
	moveLeftDy  int
	moveDownDx  int
	moveDownDy  int
	moveRightDx int
	moveRightDy int
	cellType    Cell
}

func (b Board) isEmpty(row int, col int) (bool, error) {
	n := len(b.state)
	m := len(b.state[0])
	if row >= n || row < 0 {
		return false, fmt.Errorf("row out of position %d", row)

	} else if col >= m || col < 0 {
		return false, fmt.Errorf("col out of position %d", col)
	}

	return b.state[row][col] == Empty, nil
}

func (p *Piece) move(direction Direction, board *Board) bool {

	switch direction {
	case Up:
		newRow := p.row + p.moveUpDy
		newCol := p.col + p.moveUpDy
		validMove, err := board.isEmpty(newRow, newCol)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if validMove {
			board.state[p.row][p.col] = Empty
			p.row = newRow
			p.col = newCol
			board.state[p.row][p.col] = p.cellType
		}

	}
	return true
}

func main() {
	fmt.Println("Let's play chess!")

	board := Board{}

	board.state[0][0] = Rook
	board.state[0][1] = Horse
	board.state[0][2] = Bishop
	board.state[0][3] = Queen
	board.state[0][4] = King
	board.state[0][5] = Bishop
	board.state[0][6] = Horse
	board.state[0][7] = Rook

	board.state[7][0] = Rook
	board.state[7][1] = Horse
	board.state[7][2] = Bishop
	board.state[7][4] = King
	board.state[7][3] = Queen
	board.state[7][5] = Bishop
	board.state[7][6] = Horse
	board.state[7][7] = Rook

	for row := range board.state {
		if row == 1 || row == 6 {
			for col := range board.state[row] {
				board.state[row][col] = Pawn
			}
		}
	}

	for _, row := range board.state {
		fmt.Println(row)
	}

}
