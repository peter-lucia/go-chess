package main

import (
	"fmt"
	"math"
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

func (b Board) moveIsWithinBoardBounds(row int, col int) (bool, error) {
	n := len(b.state)
	m := len(b.state[0])
	if row >= n || row < 0 {
		return false, fmt.Errorf("row out of position %d", row)

	} else if col >= m || col < 0 {
		return false, fmt.Errorf("col out of position %d", col)
	}

	return b.state[row][col] == Empty, nil
}

func (b Board) moveIsValidForPiece(movingPieceType Cell, rowDy int, colDx int) (bool, error) {

	switch movingPieceType {
	case Queen:
		if math.Abs(float64(rowDy)) == math.Abs(float64(colDx)) { // diagonal
			return true, nil
		} else if rowDy == 0 || colDx == 0 { // up / down / left / right
			return true, nil
		} else {
			return false, nil
		}
	case King:
		if rowDy == 1 && colDx == 0 || rowDy == 0 && colDx == 1 {
			return true, nil
		} else {
			return false, nil
		}
	case Bishop:
		if math.Abs(float64(rowDy)) == math.Abs(float64(colDx)) { // diagonal
			return true, nil
		}
	case Horse:
		if math.Abs(float64(rowDy)) == 2 && math.Abs(float64(colDx)) == 1 {
			return true, nil
		} else if math.Abs(float64(rowDy)) == 1 && math.Abs(float64(colDx)) == 2 {
			return true, nil

		}

	}

	return false, nil

}

func (b Board) moveReachesEmptyCellOrOpponent(movingPieceType Cell, rowDy int, colDx int) (bool, error) {
	// TODO: Implement this
	return false, nil
}

func (b Board) moveDoesNotCauseImplicitSelfCheck(movingPieceType Cell, rowDy int, colDx int) (bool, error) {
	// TODO: Implement this
	return false, nil
}

func (p *Piece) move(rowDy int, colDx int, board *Board) bool {

	newRow := p.row + rowDy
	newCol := p.col + colDx
	moveWithinBounds, err := board.moveIsWithinBoardBounds(newRow, newCol)
	moveValidForPiece, err := board.moveIsValidForPiece(p.cellType, rowDy, colDx)
	moveDestinationOK, err := board.moveReachesEmptyCellOrOpponent(p.cellType, rowDy, colDx)
	moveDoesNotCauseCheck, err := board.moveDoesNotCauseImplicitSelfCheck(p.cellType, rowDy, colDx)

	// * check that the move doesn't result in check for your king
	if err != nil {
		fmt.Println(err)
		return false
	}

	validMove := moveWithinBounds && moveValidForPiece && moveDestinationOK && moveDoesNotCauseCheck

	if validMove {
		board.state[p.row][p.col] = Empty
		p.row = newRow
		p.col = newCol
		board.state[p.row][p.col] = p.cellType
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
