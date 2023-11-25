package engine

import (
	"fmt"
	"github.com/google/uuid"
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
	State          [8][8]Piece
	BoardIsFlipped bool
	IsPlayer1Turn  bool
	Uuid           uuid.UUID
}

type Piece struct {
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
	row := 0
	col := 0
	_rowDy := rowDy
	_colDx := colDx
	for _rowDy != 0 || _colDx != 0 {
		if b.State[startRow+_rowDy][startCol+_colDx].CellType != Empty && !(row == startRow && col == startCol) && !(_rowDy == rowDy && _colDx == colDx) {
			// if we're not at the starting location or the destination location and there is a piece
			fmt.Println("Found", b.State[startRow+_rowDy][startCol+_colDx].CellType)
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
	fmt.Println("Pieces found", piecesFound)
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

func (b Board) moveReachesEmptyCellOrOpponent(p *Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {
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

	for _, row := range b.State {
		for _, col := range row {
			fmt.Printf("%02d ", col.CellType) // %02d will pad 0's to make the width 2
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

func (b Board) movePutsMovingPlayersKingInCheck(p Piece, startRow int, startCol int, rowDy int, colDx int) (bool, error) {
	/*
		Returns true if the move would put the moving player's king in check, false otherwise
	*/

	isP1, _ := p.isPlayer1()
	tempB := b
	tempB.State[startRow][startCol] = Piece{CellType: Empty}
	tempB.State[startRow+rowDy][startCol+colDx] = p
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

func (p *Piece) validMove(startRow int, startCol int, rowDy int, colDx int, board *Board) (bool, error) {

	newRow := startRow + rowDy
	newCol := startCol + colDx
	fmt.Println("Board UUID", board.Uuid.String())
	moveByCorrectPlayer, _ := board.moveByCorrectPlayer(*p)
	fmt.Println("Move by correct player", moveByCorrectPlayer)
	if !moveByCorrectPlayer {
		return false, nil
	}
	moveWithinBounds, _ := board.moveIsWithinBoardBounds(newRow, newCol)
	fmt.Println("Move within bounds", moveWithinBounds)
	if !moveWithinBounds {
		return false, nil
	}
	moveValidForPiece, _ := board.moveIsValidForPiece(*p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move valid for piece", moveValidForPiece)
	if !moveValidForPiece {
		return false, nil
	}
	moveReachesEmptyCellOrOpponent, _ := board.moveReachesEmptyCellOrOpponent(p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move reaches empty cell or opponent", moveReachesEmptyCellOrOpponent)
	if !moveReachesEmptyCellOrOpponent {
		return false, nil
	}
	moveCauseSelfCheck, _ := board.movePutsMovingPlayersKingInCheck(*p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move causes self check", moveCauseSelfCheck)
	if moveCauseSelfCheck {
		return false, nil

	}
	moveJumpsPiecesCorrectly, _ := board.moveJumpsPiecesCorrectly(*p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move jumps pieces correctly", moveJumpsPiecesCorrectly)
	if !moveJumpsPiecesCorrectly {
		return false, nil
	}
	board.printGame()
	return moveByCorrectPlayer && moveWithinBounds && moveValidForPiece && moveReachesEmptyCellOrOpponent && !moveCauseSelfCheck && moveJumpsPiecesCorrectly, nil
}

func (p *Piece) Move(startRow int, startCol int, rowDy int, colDx int, board *Board) (bool, Board, error) {

	validMove, err := p.validMove(startRow, startCol, rowDy, colDx, board)
	if err != nil {
		return false, *board, err
	}

	if validMove {
		board.State[startRow][startCol] = Piece{CellType: Empty}
		board.State[startRow+rowDy][startCol+colDx] = *p
	} else {
		return false, *board, nil
	}

	return true, *board, nil
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

func ReverseBoardRow(arr [8]Piece) ([8]Piece, error) {
	// TODO: Update piece rows and cols or remove this from the piece attribute
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr, nil
}

func FlipBoard(board Board) (Board, error) {

	board.BoardIsFlipped = !board.BoardIsFlipped

	// rotate clockwise twice
	for k := 0; k < 2; k++ {
		// pivot axes by swapping state[i][j] with state[j][i]
		for i := 0; i < len(board.State); i++ {
			for j := i; j < len(board.State[0]); j++ {
				tmp := board.State[i][j]
				board.State[i][j] = board.State[j][i]
				board.State[j][i] = tmp
			}
		}

		// reverse each row
		for i := 0; i < len(board.State); i++ {
			board.State[i], _ = ReverseBoardRow(board.State[i])
		}
	}
	return board, nil

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
	board.BoardIsFlipped = false
	board.IsPlayer1Turn = true
	board.Uuid, _ = uuid.NewUUID()

	board.State[7][0] = Piece{CellType: P2Rook}
	board.State[7][1] = Piece{CellType: P2Horse}
	board.State[7][2] = Piece{CellType: P2Bishop}
	board.State[7][3] = Piece{CellType: P2King}
	board.State[7][4] = Piece{CellType: P2Queen}
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
