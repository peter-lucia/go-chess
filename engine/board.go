package engine

import (
	"fmt"
	"github.com/google/uuid"
	"math"
)

type Board struct {
	State          [8][8]Piece
	BoardIsFlipped bool
	IsPlayer1Turn  bool
	CheckMate      bool
	Winner         string
	Uuid           uuid.UUID
}

func (b Board) createCopy() (Board, error) {
	newBoard := Board{}
	newBoard.BoardIsFlipped = b.BoardIsFlipped
	newBoard.IsPlayer1Turn = b.IsPlayer1Turn
	newBoard.CheckMate = b.CheckMate
	// Uuid should have a proper copy mechanism
	newBoard.Uuid = b.Uuid
	// board has non-primitive types so we can't copy directly
	for row, _ := range b.State {
		for col, p := range b.State[row] {
			// since Piece only has primitive types, we can copy it like this
			newBoard.State[row][col] = p
		}
	}
	return newBoard, nil
}

func (b Board) moveIsWithinBoardBounds(row int, col int) (bool, error) {
	n := len(b.State)
	m := len(b.State[0])
	if row >= n || row < 0 {
		return false, fmt.Errorf("Row out of position %d", row)

	} else if col >= m || col < 0 {
		return false, fmt.Errorf("Col out of position %d", col)
	}

	return true, nil
}

func (b Board) piecesBetweenOriginAndDestination(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {
	if p.CellType == Empty || p.CellType == P1Horse || p.CellType == P2Horse {
		return false, nil
	}

	piecesFound := 0
	_rowDy := rowDy
	_colDx := colDx
	for _rowDy != 0 || _colDx != 0 {
		if b.State[startRow+_rowDy][startCol+_colDx].CellType != Empty && !((startRow+_rowDy) == startRow && (startCol+_colDx) == startCol) && !(_rowDy == rowDy && _colDx == colDx) {
			// if we're not at the starting location or the destination location and there is a piece
			piecesFound++
		}
		if _rowDy > 0 {
			_rowDy--
		} else if _rowDy < 0 {
			_rowDy++
		}
		if _colDx > 0 {
			_colDx--
		} else if _colDx < 0 {
			_colDx++
		}
	}
	return piecesFound > 0, nil
}

func (b Board) moveJumpsPiecesCorrectly(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {

	// Check if a non horse piece tries to jump other pieces between
	// the origin and destination
	if p.CellType != P1Horse && p.CellType != P2Horse {
		piecesBetweenOrigDest, _ := b.piecesBetweenOriginAndDestination(p, startRow, startCol, rowDy, colDx)
		if piecesBetweenOrigDest {
			return false, nil
		}
	}
	return true, nil
}

func (b Board) moveIsValidForPiece(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {

	if (startRow+rowDy < 0) || (startRow+rowDy > 7) {
		return false, nil
	}
	if (startCol+colDx < 0) || (startCol+colDx > 7) {
		return false, nil
	}

	isP1, _ := p.isPlayer1()
	switch p.CellType {
	case P1Queen, P2Queen:
		if math.Abs(float64(rowDy)) == math.Abs(float64(colDx)) { // diagonal
			return true, nil
		} else if rowDy == 0 || colDx == 0 { // up / down / left / right
			return true, nil
		} else {
			return false, nil
		}
	case P1King, P2King:
		if math.Abs(float64(rowDy)) == 1 && colDx == 0 || rowDy == 0 && math.Abs(float64(colDx)) == 1 || math.Abs(float64(rowDy)) == 1 && math.Abs(float64(colDx)) == 1 {
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
	case P1Rook, P2Rook:
		if math.Abs(float64(rowDy)) > 0 && colDx == 0 || rowDy == 0 && math.Abs(float64(colDx)) > 0 {
			return true, nil
		}
	case P1Pawn, P2Pawn:
		if isP1 && !b.BoardIsFlipped && rowDy == 1 && math.Abs(float64(colDx)) <= 1 {
			return true, nil
		} else if isP1 && b.BoardIsFlipped && rowDy == -1 && math.Abs(float64(colDx)) <= 1 {
			return true, nil
		} else if !isP1 && !b.BoardIsFlipped && rowDy == -1 && math.Abs(float64(colDx)) <= 1 {
			return true, nil
		} else if !isP1 && b.BoardIsFlipped && rowDy == 1 && math.Abs(float64(colDx)) <= 1 {
			return true, nil
		} else if math.Abs(float64(rowDy)) == 2 && (startRow == 1 || startRow == 6) {
			return true, nil
		}
	}

	return false, nil

}

func (b Board) moveReachesEmptyCellOrOpponent(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {
	rows := len(b.State)
	cols := len(b.State[0])
	newRow := startRow + rowDy
	newCol := startCol + colDx

	if newRow < 0 || newRow >= rows {
		return false, nil
	}

	if newCol < 0 || newCol >= cols {
		return false, nil
	}

	if p.CellType == Empty {
		return false, fmt.Errorf("a piece cannot have an empty cell type")
	}

	destinationPiece := b.State[newRow][newCol]
	if destinationPiece.CellType != Empty {
		// check if destination piece is owned by the moving player
		movingPieceIsP1, _ := p.isPlayer1()
		destinationPieceIsP1, _ := destinationPiece.isPlayer1()
		if movingPieceIsP1 == destinationPieceIsP1 {
			return false, nil
		}
		if (p.CellType == P1Pawn || p.CellType == P2Pawn) && (colDx == 0 && math.Abs(float64(rowDy)) >= 1) {
			// pawns can't kill an opponent directly in front of them
			return false, nil
		}
	} else if (p.CellType == P1Pawn || p.CellType == P2Pawn) && math.Abs(float64(colDx)) > 0 {
		// pawns cannot go diagonally to an empty space
		return false, nil
	}

	return true, nil
}

func (b Board) printGame() {

	if b.BoardIsFlipped {
		fmt.Println("        Player 1")
	} else {
		fmt.Println("        Player 2")
	}

	for row := len(b.State) - 1; row >= 0; row-- {
		for col := 0; col < len(b.State[row]); col++ {
			fmt.Printf("%02d ", b.State[row][col].CellType) // %02d will pad 0's to make the width 2
		}
		fmt.Println()
	}
	if b.BoardIsFlipped {
		fmt.Println("        Player 2")
	} else {
		fmt.Println("        Player 1")
	}
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
		for col, p := range b.State[row] {
			isP1, _ := p.isPlayer1()
			hasKingInCheck, _ := p.hasKingInCheck(b, row, col)
			if isP1 && hasKingInCheck {
				p2KingInCheck = true
			} else if hasKingInCheck {
				p1KingInCheck = true
			}
		}
	}

	return p1KingInCheck, p2KingInCheck, nil

}

func (b *Board) GetPositionsOfKings() (int, int, int, int, error) {
	p1KingRow := -1
	p1KingCol := -1
	p2KingRow := -1
	p2KingCol := -1
	for row, _ := range b.State {
		for col, p := range b.State[row] {
			if p.CellType == P1King {
				p1KingRow = row
				p1KingCol = col
			} else if p.CellType == P2King {
				p2KingRow = row
				p2KingCol = col
			}
		}
	}

	return p1KingRow, p1KingCol, p2KingRow, p2KingCol, nil

}

func (b Board) isCheckmate() (bool, string, error) {
	p1KingInCheck, p2KingInCheck, err := b.isEitherKingInCheck()
	if err != nil {
		return false, "none", err
	}

	if !p1KingInCheck && !p2KingInCheck {
		return false, "none", nil
	}

	// if king is in check, he must have at least one move that doesn't put him in check
	// check all 8 possible moves for the king to see if they are valid
	p1KingRow, p1KingCol, p2KingRow, p2KingCol, err := b.GetPositionsOfKings()

	fmt.Println("P1KingRow", p1KingRow, "P1KingCol", p1KingCol)
	fmt.Println("P2KingRow", p2KingRow, "P2KingCol", p2KingCol)

	kingPiece := b.State[p1KingRow][p1KingCol]
	possibleWinner := "Player 2"

	kingRow := p1KingRow
	kingCol := p1KingCol
	if p2KingInCheck {
		kingPiece = b.State[p2KingRow][p2KingCol]
		possibleWinner = "Player 1"
		kingRow = p2KingRow
		kingCol = p2KingCol
	}
	kingPieceIsP1, _ := kingPiece.isPlayer1()
	fmt.Println("KingPiece is Player 1?", kingPieceIsP1)

	// TODO: check if any other piece can get in the way of putting the king in check
	// TODO: check if any other piece can kill the opponent's piece putting the king in check

	// If no other piece can intervene, check if the king can move
	for colDx := -1; colDx <= 1; colDx++ {
		for rowDy := -1; rowDy <= 1; rowDy++ {
			if colDx == 0 && rowDy == 0 {
				continue
			}
			bTemp, _ := b.createCopy()

			moveValidForPiece, _ := bTemp.moveIsValidForPiece(kingPiece, kingRow, kingCol, rowDy, colDx)
			if !moveValidForPiece {
				continue
			}

			moveToValidPosition, _ := bTemp.moveReachesEmptyCellOrOpponent(kingPiece, kingRow, kingCol, rowDy, colDx)
			if !moveToValidPosition {
				continue
			}
			putsKingInCheck, _ := bTemp.movePutsMovingPlayersKingInCheck(kingPiece, kingRow, kingCol, rowDy, colDx)
			if !putsKingInCheck {
				return false, "none", nil

			}

		}
	}

	// checkmate
	return true, possibleWinner, nil

}

func (b Board) isStalemate() (bool, error) {
	// if king is not in check, but is the only piece left, he must be able to move somewhere
	// without being put in check

	return false, nil
}

func (b Board) movePutsMovingPlayersKingInCheck(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {
	/*
		Returns true if the move would put the moving player's king in check, false otherwise
	*/

	isP1, _ := p.isPlayer1()
	tempB, _ := b.createCopy()
	tempB.State[startRow][startCol] = Piece{CellType: Empty}
	tempB.State[startRow+rowDy][startCol+colDx] = Piece{CellType: p.CellType}
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

func (b Board) moveByCorrectPlayer(p Piece) (bool, error) {
	isP1, _ := p.isPlayer1()
	return (b.IsPlayer1Turn && isP1) || (!b.IsPlayer1Turn && !isP1), nil
}

func (b *Board) convertPawns(convertToCellType Cell) (bool, error) {
	for col := range b.State[0] {
		if b.State[0][col].CellType == P2Pawn {
			b.State[0][col].CellType = convertToCellType
		}
	}

	for col := range b.State[7] {
		if b.State[7][col].CellType == P1Pawn {
			b.State[7][col].CellType = convertToCellType
		}
	}
	return true, nil

}
