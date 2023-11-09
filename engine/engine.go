package engine

import (
	"fmt"
	"math"
)

type Cell int

const (
	Empty Cell = iota

	P1King
	P1Queen
	P1Bishop
	P1Horse
	P1Rook
	P1Pawn

	P2King
	P2Queen
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
	State [8][8]Piece
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
	CellType    Cell
}

func (p Piece) hasKingInCheck(b Board) (bool, error) {
	// Go through all the piece's moves and check if any one of them is valid
	// and could kill the opposing player's king
	return false, nil
}

func (p Piece) isPlayer1() (bool, error) {
	return p.CellType > Empty && p.CellType <= P1Pawn, nil
}

func (b Board) moveIsWithinBoardBounds(row int, col int) (bool, error) {
	n := len(b.State)
	m := len(b.State[0])
	if row >= n || row < 0 {
		return false, fmt.Errorf("row out of position %d", row)

	} else if col >= m || col < 0 {
		return false, fmt.Errorf("col out of position %d", col)
	}

	return b.State[row][col].CellType == Empty, nil
}

func (b Board) moveIsValidForPiece(p Piece, rowDy int, colDx int) (bool, error) {

	movingPieceType := p.CellType

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
	case P1Pawn:
		if rowDy == 1 || (rowDy == 2 && p.row == 1) {
			return true, nil
		}
	case P2Pawn:
		if rowDy == -1 || (rowDy == -2 && p.row == 6) {
			return true, nil
		}
	}

	return false, nil

}

func (b Board) moveReachesEmptyCellOrOpponent(p *Piece, rowDy int, colDx int) (bool, error) {
	rows := len(b.State)
	cols := len(b.State[0])
	newRow := p.row + rowDy
	newCol := p.col + colDx

	if newRow < 0 || newRow >= rows {
		return false, nil
	}

	if newCol < 0 || newCol >= cols {
		return false, nil
	}

	if p.CellType == Empty {
		return false, fmt.Errorf("a piece cannot have an empty cell type")
	}

	destinationCell := b.State[newRow][newCol].CellType
	// player 1 trying to move to existing p1 occupied cell
	if p.CellType <= P1Pawn && destinationCell > Empty && destinationCell <= P1Pawn {
		return false, nil
	}

	// player 2 trying to move to existing p2 occupied cell
	if p.CellType > P1Pawn && destinationCell > P1Pawn {
		return false, nil
	}

	return true, nil
}

func (b Board) getCoordsForCellType(cellType Cell) ([2]int, error) {

	for rowNum, row := range b.State {
		for colNum, p := range row {
			if p.CellType == cellType {
				return [2]int{rowNum, colNum}, nil
			}
		}
	}
	return [2]int{-1, -1}, nil
}

func (b Board) isEitherKingInCheck() (bool, bool, error) {
	p1KingInCheck := false
	p2KingInCheck := false
	for row, _ := range b.State {
		for _, p := range b.State[row] {
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

func (b Board) moveDoesNotPutPieceKingInCheck(p Piece, rowDy int, colDx int) (bool, error) {
	/*
		Returns true if the move would put the moving player's king in check, false otherwise
	*/

	isP1, _ := p.isPlayer1()
	tempB := b
	tempB.State[p.row][p.col] = Piece{CellType: Empty}
	tempB.State[p.row+rowDy][p.col+colDx] = p
	p1InCheck, p2InCheck, _ := tempB.isEitherKingInCheck()

	if p1InCheck && isP1 {
		return true, nil
	}

	if p2InCheck && !isP1 {
		return true, nil
	}

	return false, nil

}

func (b Board) moveDoesNotIgnoreCurrentKingInCheck(p *Piece, rowDy int, colDx int) (bool, error) {

	return false, nil

}

func (p *Piece) validMove(rowDy int, colDx int, board *Board) (bool, error) {

	newRow := p.row + rowDy
	newCol := p.col + colDx
	moveWithinBounds, _ := board.moveIsWithinBoardBounds(newRow, newCol)
	moveValidForPiece, _ := board.moveIsValidForPiece(*p, rowDy, colDx)
	moveDestinationOK, _ := board.moveReachesEmptyCellOrOpponent(p, rowDy, colDx)
	moveDoesNotCauseCheck, _ := board.moveDoesNotPutPieceKingInCheck(*p, rowDy, colDx)
	return moveWithinBounds && moveValidForPiece && moveDestinationOK && moveDoesNotCauseCheck, nil
}

func (p *Piece) move(rowDy int, colDx int, board *Board) (bool, error) {

	validMove, err := p.validMove(rowDy, colDx, board)
	if err != nil {
		return false, err
	}

	newRow := p.row + rowDy
	newCol := p.col + colDx
	if validMove {
		board.State[p.row][p.col] = Piece{CellType: Empty}
		p.row = newRow
		p.col = newCol
		board.State[p.row][p.col] = *p
	}

	return true, nil
}

func InitGame() (Board, error) {

	board := Board{}

	board.State[0][0] = Piece{CellType: P1Rook}
	board.State[0][1] = Piece{CellType: P1Horse}
	board.State[0][2] = Piece{CellType: P1Bishop}
	board.State[0][3] = Piece{CellType: P1Queen}
	board.State[0][4] = Piece{CellType: P1King}
	board.State[0][5] = Piece{CellType: P1Bishop}
	board.State[0][6] = Piece{CellType: P1Horse}
	board.State[0][7] = Piece{CellType: P1Rook}

	board.State[7][0] = Piece{CellType: P2Rook}
	board.State[7][1] = Piece{CellType: P2Horse}
	board.State[7][2] = Piece{CellType: P2Bishop}
	board.State[7][4] = Piece{CellType: P2King}
	board.State[7][3] = Piece{CellType: P2Queen}
	board.State[7][5] = Piece{CellType: P2Bishop}
	board.State[7][6] = Piece{CellType: P2Horse}
	board.State[7][7] = Piece{CellType: P2Rook}

	for row := range board.State {
		for col := range board.State[row] {
			if row == 1 {
				board.State[row][col] = Piece{CellType: P1Pawn}
			} else if row == 6 {
				board.State[row][col] = Piece{CellType: P2Pawn}
			}
		}
	}

	return board, nil

}
