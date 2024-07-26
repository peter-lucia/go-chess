package engine

import "fmt"

type Piece struct {
	// TODO: Use these, convert to vectors for non-jumping pieces, or get rid of them
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

func (p Piece) validMove(startRow int, startCol int, rowDy int, colDx int, board *Board) (bool, error) {
	/*
		Check if the provided move parameters constitutes a valid move
	*/

	newRow := startRow + rowDy
	newCol := startCol + colDx
	fmt.Println("Board UUID", board.Uuid.String())
	moveByCorrectPlayer, _ := board.moveByCorrectPlayer(p)
	fmt.Println("Move by correct player", moveByCorrectPlayer)
	if !moveByCorrectPlayer {
		return false, nil
	}
	moveWithinBounds, _ := board.moveIsWithinBoardBounds(newRow, newCol)
	fmt.Println("Move within bounds", moveWithinBounds)
	if !moveWithinBounds {
		return false, nil
	}
	moveValidForPiece, _ := board.moveIsValidForPiece(p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move valid for piece", moveValidForPiece)
	if !moveValidForPiece {
		return false, nil
	}
	moveReachesEmptyCellOrOpponent, _ := board.moveReachesEmptyCellOrOpponent(p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move reaches empty cell or opponent", moveReachesEmptyCellOrOpponent)
	if !moveReachesEmptyCellOrOpponent {
		return false, nil
	}
	moveCauseSelfCheck, _ := board.movePutsMovingPlayersKingInCheck(p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move causes self check", moveCauseSelfCheck)
	if moveCauseSelfCheck {
		return false, nil

	}
	moveJumpsPiecesCorrectly, _ := board.moveJumpsPiecesCorrectly(p, startRow, startCol, rowDy, colDx)
	fmt.Println("Move jumps pieces correctly", moveJumpsPiecesCorrectly)
	if !moveJumpsPiecesCorrectly {
		return false, nil
	}

	board.printGame()
	return moveByCorrectPlayer && moveWithinBounds && moveValidForPiece && moveReachesEmptyCellOrOpponent && !moveCauseSelfCheck && moveJumpsPiecesCorrectly, nil
}

func (p *Piece) Move(startRow int, startCol int, rowDy int, colDx int, board *Board) (bool, Board, error) {
	/*
		Move the piece on the chess board
	*/

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

func (p Piece) hasKingInCheck(b Board, pieceRow int, pieceCol int) (bool, error) {

	// See if moving the piece to the king's position is a valid move for the piece
	p1KingRow, p1KingCol, p2KingRow, p2KingCol, _ := b.GetPositionsOfKings()

	isP1, _ := p.isPlayer1()

	rowDy := p2KingRow - pieceRow
	colDx := p2KingCol - pieceCol
	if !isP1 {
		rowDy = p1KingRow - pieceRow
		colDx = p1KingCol - pieceCol
	}
	validForPiece, _ := b.moveIsValidForPiece(p, pieceRow, pieceCol, rowDy, colDx)
	moveReachesEmptyCellOrOpponent, _ := b.moveReachesEmptyCellOrOpponent(p, pieceRow, pieceCol, rowDy, colDx)
	moveJumpsPiecesCorrectly, _ := b.moveJumpsPiecesCorrectly(p, pieceRow, pieceCol, rowDy, colDx)
	return validForPiece && moveReachesEmptyCellOrOpponent && moveJumpsPiecesCorrectly, nil

}

func (p Piece) isPlayer1() (bool, error) {
	return p.CellType > Empty && p.CellType <= P1Pawn, nil
}
