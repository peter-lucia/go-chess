package main

import (
	"fmt"
	"math"
)

type Cell int

const (
	Empty Cell = iota

	P1Queen
	P1King
	P1Bishop
	P1Horse
	P1Rook
	P1Pawn

	P2Queen
	P2King
	P2Bishop
	P2Horse
	P2Rook
	P2Pawn
	// be sure to update moveReachesEmptyCellOrOpponent before adding values here

)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Board struct {
	state [8][8]Piece
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

func (p Piece) hasKingInCheck(b Board) (bool, error) {
	// Go through all of the piece's moves and check if any one of them is valid
	// and could kill the opposing player's king
	return false, nil
}

func (p Piece) isPlayer1() (bool, error) {
	return p.cellType > Empty && p.cellType <= P1Pawn, nil
}

func (b Board) moveIsWithinBoardBounds(row int, col int) (bool, error) {
	n := len(b.state)
	m := len(b.state[0])
	if row >= n || row < 0 {
		return false, fmt.Errorf("row out of position %d", row)

	} else if col >= m || col < 0 {
		return false, fmt.Errorf("col out of position %d", col)
	}

	return b.state[row][col].cellType == Empty, nil
}

func (b Board) moveIsValidForPiece(movingPieceType Cell, rowDy int, colDx int) (bool, error) {

	switch movingPieceType {
	case P1Queen, P2Queen:
		if math.Abs(float64(rowDy)) == math.Abs(float64(colDx)) { // diagonal
			return true, nil
		} else if rowDy == 0 || colDx == 0 { // up / down / left / right
			return true, nil
		} else {
			return false, nil
		}
	case P1King, P2King:
		if rowDy == 1 && colDx == 0 || rowDy == 0 && colDx == 1 {
			return true, nil
		} else {
			return false, nil
		}
	case P1Bishop, P2Bishop:
		if math.Abs(float64(rowDy)) == math.Abs(float64(colDx)) { // diagonal
			return true, nil
		}
	case P1Horse, P2Horse:
		if math.Abs(float64(rowDy)) == 2 && math.Abs(float64(colDx)) == 1 {
			return true, nil
		} else if math.Abs(float64(rowDy)) == 1 && math.Abs(float64(colDx)) == 2 {
			return true, nil

		}
	}

	return false, nil

}

func (b Board) moveReachesEmptyCellOrOpponent(p *Piece, rowDy int, colDx int) (bool, error) {
	rows := len(b.state)
	cols := len(b.state[0])
	newRow := p.row + rowDy
	newCol := p.col + colDx

	if newRow < 0 || newRow >= rows {
		return false, nil
	}

	if newCol < 0 || newCol >= cols {
		return false, nil
	}

	if p.cellType == Empty {
		return false, fmt.Errorf("a piece cannot have an empty cell type")
	}

	destinationCell := b.state[newRow][newCol].cellType
	// player 1 trying to move to existing p1 occupied cell
	if p.cellType <= P1Pawn && destinationCell > Empty && destinationCell <= P1Pawn {
		return false, nil
	}

	// player 2 trying to move to existing p2 occupied cell
	if p.cellType > P1Pawn && destinationCell > P1Pawn {
		return false, nil
	}

	return true, nil
}

func (b Board) getCoordsForCellType(cellType Cell) ([2]int, error) {

	for rowNum, row := range b.state {
		for colNum, p := range row {
			if p.cellType == cellType {
				return [2]int{rowNum, colNum}, nil
			}
		}
	}
	return [2]int{-1, -1}, nil
}

func (b Board) isEitherKingInCheck() (bool, bool, error) {
	p1KingInCheck := false
	p2KingInCheck := false
	for row, _ := range b.state {
		for _, p := range b.state[row] {
			isP1, _ := p.isPlayer1()
			hasKingInCheck, _ := p.hasKingInCheck(b)
			if isP1 && hasKingInCheck {
				p2KingInCheck = true
			} else if hasKingInCheck {
				p1KingInCheck = true
			}
		}
	}

	return p1KingInCheck, p2KingInCheck, nil

}

func (b Board) moveDoesNotPutPlayersKingInCheck(p Piece, rowDy int, colDx int) (bool, error) {

	isP1, _ := p.isPlayer1()
	tempB := b
	tempB.state[p.row][p.col] = Piece{cellType: Empty}
	tempB.state[p.row+rowDy][p.col+colDx] = p
	p1InCheck, _, _ := tempB.isEitherKingInCheck()

	if p1InCheck && isP1 {
		return true, nil
	}

	return false, nil

}

func (b Board) moveDoesNotIgnoreCurrentKingInCheck(p *Piece, rowDy int, colDx int) (bool, error) {

	return false, nil

}

func (p *Piece) move(rowDy int, colDx int, board *Board) bool {

	newRow := p.row + rowDy
	newCol := p.col + colDx
	moveWithinBounds, err := board.moveIsWithinBoardBounds(newRow, newCol)
	moveValidForPiece, err := board.moveIsValidForPiece(p.cellType, rowDy, colDx)
	moveDestinationOK, err := board.moveReachesEmptyCellOrOpponent(p, rowDy, colDx)
	moveDoesNotCauseCheck, err := board.moveDoesNotPutPlayersKingInCheck(*p, rowDy, colDx)

	// * check that the move doesn't result in check for your king
	if err != nil {
		fmt.Println(err)
		return false
	}

	validMove := moveWithinBounds && moveValidForPiece && moveDestinationOK && moveDoesNotCauseCheck

	if validMove {
		board.state[p.row][p.col] = Piece{cellType: Empty}
		p.row = newRow
		p.col = newCol
		board.state[p.row][p.col] = *p
	}

	return true
}

func main() {
	fmt.Println("Let's play chess!")

	board := Board{}

	board.state[0][0] = Piece{cellType: P1Rook}
	board.state[0][1] = Piece{cellType: P1Horse}
	board.state[0][2] = Piece{cellType: P1Bishop}
	board.state[0][3] = Piece{cellType: P1Queen}
	board.state[0][4] = Piece{cellType: P1King}
	board.state[0][5] = Piece{cellType: P1Bishop}
	board.state[0][6] = Piece{cellType: P1Horse}
	board.state[0][7] = Piece{cellType: P1Rook}

	board.state[7][0] = Piece{cellType: P1Rook}
	board.state[7][1] = Piece{cellType: P1Horse}
	board.state[7][2] = Piece{cellType: P1Bishop}
	board.state[7][4] = Piece{cellType: P1King}
	board.state[7][3] = Piece{cellType: P1Queen}
	board.state[7][5] = Piece{cellType: P1Bishop}
	board.state[7][6] = Piece{cellType: P1Horse}
	board.state[7][7] = Piece{cellType: P1Rook}

	for row := range board.state {
		if row == 1 || row == 6 {
			for col := range board.state[row] {
				board.state[row][col] = Piece{cellType: P1Pawn}
			}
		}
	}

	for _, row := range board.state {
		for _, col := range row {
			fmt.Printf("%d ", col.cellType)
		}
		fmt.Println()

	}

}
