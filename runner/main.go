package main

import (
	"errors"
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
	"github.com/peter-lucia/go-chess/ui"
)

// map UUID to a board object pointer
var gameLookup = make(map[string]*engine.Board)

func convertToUIPiece(piece engine.Piece) (string, error) {
	// {wR wN wB wQ wK wB wN wR wP wP wP wP wP wP wP wP ... bP bP bP bP bP bP bP bP bR bN bB bQ bK bB bN bR}

	switch piece.CellType {
	case engine.P1Rook:
		return "wR", nil
	case engine.P1Horse:
		return "wN", nil
	case engine.P1Bishop:
		return "wB", nil
	case engine.P1Queen:
		return "wQ", nil
	case engine.P1King:
		return "wK", nil
	case engine.P1Pawn:
		return "wP", nil

	case engine.P2Rook:
		return "bR", nil
	case engine.P2Horse:
		return "bN", nil
	case engine.P2Bishop:
		return "bB", nil
	case engine.P2Queen:
		return "bQ", nil
	case engine.P2King:
		return "bK", nil
	case engine.P2Pawn:
		return "bP", nil
	default:
		return "empty", nil
	}

}

func convertUICoordsToEngineCoords(uiPiece string) (int, int, error) {
	if len(uiPiece) != 2 {
		return 0, 0, errors.New("uninterpretable ui coords")
	}
	// cols: a,b,c,d,e,f,g,h
	col := int(uiPiece[0] - 'a')
	// rows: 1,2,3,4,5,6,7
	row := int(uiPiece[1] - '1')
	return row, col, nil

}

func convertToEnginePieceOnBoard(uiPiece string, row int, col int, b *engine.Board) (engine.Board, error) {
	// {wR wN wB wQ wK wB wN wR wP wP wP wP wP wP wP wP ... bP bP bP bP bP bP bP bP bR bN bB bQ bK bB bN bR}

	switch uiPiece {
	case "wR":
		b.State[row][col] = engine.Piece{CellType: engine.P1Rook, Row: row, Col: col}
	case "wN":
		b.State[row][col] = engine.Piece{CellType: engine.P1Horse, Row: row, Col: col}
	case "wB":
		b.State[row][col] = engine.Piece{CellType: engine.P1Bishop, Row: row, Col: col}
	case "wQ":
		b.State[row][col] = engine.Piece{CellType: engine.P1Queen, Row: row, Col: col}
	case "wK":
		b.State[row][col] = engine.Piece{CellType: engine.P1King, Row: row, Col: col}
	case "wP":
		b.State[row][col] = engine.Piece{CellType: engine.P1Pawn, Row: row, Col: col}

	case "bR":
		b.State[row][col] = engine.Piece{CellType: engine.P2Rook, Row: row, Col: col}
	case "bN":
		b.State[row][col] = engine.Piece{CellType: engine.P2Horse, Row: row, Col: col}
	case "bB":
		b.State[row][col] = engine.Piece{CellType: engine.P2Bishop, Row: row, Col: col}
	case "bQ":
		b.State[row][col] = engine.Piece{CellType: engine.P2Queen, Row: row, Col: col}
	case "bK":
		b.State[row][col] = engine.Piece{CellType: engine.P2King, Row: row, Col: col}
	case "bP":
		b.State[row][col] = engine.Piece{CellType: engine.P2Pawn, Row: row, Col: col}
	default:
		b.State[row][col] = engine.Piece{CellType: engine.Empty, Row: row, Col: col}

	}

	return *b, nil
}

func translateToEngineBoardPosition(uiBoard ui.BoardPosition) (engine.Board, error) {

	b_ptr, found := gameLookup[uiBoard.UUID]
	b := *b_ptr
	if !found {
		fmt.Println("Cannot find board")
		return b, errors.New("cannot find board")
	}
	b, _ = convertToEnginePieceOnBoard(uiBoard.A8, 7, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B8, 7, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C8, 7, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D8, 7, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E8, 7, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F8, 7, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G8, 7, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H8, 7, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A7, 6, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B7, 6, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C7, 6, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D7, 6, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E7, 6, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F7, 6, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G7, 6, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H7, 6, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A6, 5, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B6, 5, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C6, 5, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D6, 5, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E6, 5, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F6, 5, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G6, 5, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H6, 5, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A5, 4, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B5, 4, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C5, 4, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D5, 4, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E5, 4, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F5, 4, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G5, 4, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H5, 4, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A4, 3, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B4, 3, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C4, 3, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D4, 3, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E4, 3, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F4, 3, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G4, 3, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H4, 3, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A3, 2, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B3, 2, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C3, 2, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D3, 2, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E3, 2, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F3, 2, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G3, 2, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H3, 2, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A2, 1, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B2, 1, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C2, 1, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D2, 1, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E2, 1, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F2, 1, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G2, 1, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H2, 1, 7, &b)

	b, _ = convertToEnginePieceOnBoard(uiBoard.A1, 0, 0, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.B1, 0, 1, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.C1, 0, 2, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.D1, 0, 3, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.E1, 0, 4, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.F1, 0, 5, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.G1, 0, 6, &b)
	b, _ = convertToEnginePieceOnBoard(uiBoard.H1, 0, 7, &b)

	return b, nil

}

func translateToUIBoardPosition(engineBoardPosition engine.Board) (ui.BoardPosition, error) {

	uiBoard := ui.BoardPosition{}
	uiBoard.UUID = engineBoardPosition.Uuid.String()
	uiBoard.IsPlayer1Turn = engineBoardPosition.IsPlayer1Turn
	uiBoard.A8, _ = convertToUIPiece(engineBoardPosition.State[7][0])
	uiBoard.B8, _ = convertToUIPiece(engineBoardPosition.State[7][1])
	uiBoard.C8, _ = convertToUIPiece(engineBoardPosition.State[7][2])
	uiBoard.D8, _ = convertToUIPiece(engineBoardPosition.State[7][3])
	uiBoard.E8, _ = convertToUIPiece(engineBoardPosition.State[7][4])
	uiBoard.F8, _ = convertToUIPiece(engineBoardPosition.State[7][5])
	uiBoard.G8, _ = convertToUIPiece(engineBoardPosition.State[7][6])
	uiBoard.H8, _ = convertToUIPiece(engineBoardPosition.State[7][7])

	uiBoard.A7, _ = convertToUIPiece(engineBoardPosition.State[6][0])
	uiBoard.B7, _ = convertToUIPiece(engineBoardPosition.State[6][1])
	uiBoard.C7, _ = convertToUIPiece(engineBoardPosition.State[6][2])
	uiBoard.D7, _ = convertToUIPiece(engineBoardPosition.State[6][3])
	uiBoard.E7, _ = convertToUIPiece(engineBoardPosition.State[6][4])
	uiBoard.F7, _ = convertToUIPiece(engineBoardPosition.State[6][5])
	uiBoard.G7, _ = convertToUIPiece(engineBoardPosition.State[6][6])
	uiBoard.H7, _ = convertToUIPiece(engineBoardPosition.State[6][7])

	uiBoard.A6, _ = convertToUIPiece(engineBoardPosition.State[5][0])
	uiBoard.B6, _ = convertToUIPiece(engineBoardPosition.State[5][1])
	uiBoard.C6, _ = convertToUIPiece(engineBoardPosition.State[5][2])
	uiBoard.D6, _ = convertToUIPiece(engineBoardPosition.State[5][3])
	uiBoard.E6, _ = convertToUIPiece(engineBoardPosition.State[5][4])
	uiBoard.F6, _ = convertToUIPiece(engineBoardPosition.State[5][5])
	uiBoard.G6, _ = convertToUIPiece(engineBoardPosition.State[5][6])
	uiBoard.H6, _ = convertToUIPiece(engineBoardPosition.State[5][7])

	uiBoard.A5, _ = convertToUIPiece(engineBoardPosition.State[4][0])
	uiBoard.B5, _ = convertToUIPiece(engineBoardPosition.State[4][1])
	uiBoard.C5, _ = convertToUIPiece(engineBoardPosition.State[4][2])
	uiBoard.D5, _ = convertToUIPiece(engineBoardPosition.State[4][3])
	uiBoard.E5, _ = convertToUIPiece(engineBoardPosition.State[4][4])
	uiBoard.F5, _ = convertToUIPiece(engineBoardPosition.State[4][5])
	uiBoard.G5, _ = convertToUIPiece(engineBoardPosition.State[4][6])
	uiBoard.H5, _ = convertToUIPiece(engineBoardPosition.State[4][7])

	uiBoard.A4, _ = convertToUIPiece(engineBoardPosition.State[3][0])
	uiBoard.B4, _ = convertToUIPiece(engineBoardPosition.State[3][1])
	uiBoard.C4, _ = convertToUIPiece(engineBoardPosition.State[3][2])
	uiBoard.D4, _ = convertToUIPiece(engineBoardPosition.State[3][3])
	uiBoard.E4, _ = convertToUIPiece(engineBoardPosition.State[3][4])
	uiBoard.F4, _ = convertToUIPiece(engineBoardPosition.State[3][5])
	uiBoard.G4, _ = convertToUIPiece(engineBoardPosition.State[3][6])
	uiBoard.H4, _ = convertToUIPiece(engineBoardPosition.State[3][7])

	uiBoard.A3, _ = convertToUIPiece(engineBoardPosition.State[2][0])
	uiBoard.B3, _ = convertToUIPiece(engineBoardPosition.State[2][1])
	uiBoard.C3, _ = convertToUIPiece(engineBoardPosition.State[2][2])
	uiBoard.D3, _ = convertToUIPiece(engineBoardPosition.State[2][3])
	uiBoard.E3, _ = convertToUIPiece(engineBoardPosition.State[2][4])
	uiBoard.F3, _ = convertToUIPiece(engineBoardPosition.State[2][5])
	uiBoard.G3, _ = convertToUIPiece(engineBoardPosition.State[2][6])
	uiBoard.H3, _ = convertToUIPiece(engineBoardPosition.State[2][7])

	uiBoard.A2, _ = convertToUIPiece(engineBoardPosition.State[1][0])
	uiBoard.B2, _ = convertToUIPiece(engineBoardPosition.State[1][1])
	uiBoard.C2, _ = convertToUIPiece(engineBoardPosition.State[1][2])
	uiBoard.D2, _ = convertToUIPiece(engineBoardPosition.State[1][3])
	uiBoard.E2, _ = convertToUIPiece(engineBoardPosition.State[1][4])
	uiBoard.F2, _ = convertToUIPiece(engineBoardPosition.State[1][5])
	uiBoard.G2, _ = convertToUIPiece(engineBoardPosition.State[1][6])
	uiBoard.H2, _ = convertToUIPiece(engineBoardPosition.State[1][7])

	uiBoard.A1, _ = convertToUIPiece(engineBoardPosition.State[0][0])
	uiBoard.B1, _ = convertToUIPiece(engineBoardPosition.State[0][1])
	uiBoard.C1, _ = convertToUIPiece(engineBoardPosition.State[0][2])
	uiBoard.D1, _ = convertToUIPiece(engineBoardPosition.State[0][3])
	uiBoard.E1, _ = convertToUIPiece(engineBoardPosition.State[0][4])
	uiBoard.F1, _ = convertToUIPiece(engineBoardPosition.State[0][5])
	uiBoard.G1, _ = convertToUIPiece(engineBoardPosition.State[0][6])
	uiBoard.H1, _ = convertToUIPiece(engineBoardPosition.State[0][7])

	return uiBoard, nil

}

func handleFlip(mr ui.RequestFlip) (bool, ui.BoardPosition, error) {
	engineBoardPosition, _ := translateToEngineBoardPosition(mr.CurrentBoardPosition)

	engineBoardPosition, _ = engine.FlipBoard(engineBoardPosition)
	gameLookup[engineBoardPosition.Uuid.String()] = &engineBoardPosition
	uiBoardPosition, _ := translateToUIBoardPosition(engineBoardPosition)
	return true, uiBoardPosition, nil

}

func handleInit() (bool, ui.BoardPosition, error) {
	engineNewBoardPosition, _ := engine.InitGame()
	gameLookup[engineNewBoardPosition.Uuid.String()] = &engineNewBoardPosition
	uiNewBoardPosition, _ := translateToUIBoardPosition(engineNewBoardPosition)
	return true, uiNewBoardPosition, nil

}

func handleMove(mr ui.RequestMove) (bool, ui.BoardPosition, error) {
	// Returns true, new board detail, nil if the move was successful
	// returns false, new board detail, nil if the move was a failure
	// returns an error if there was a problem with the move
	fmt.Println("Start", mr.Start, "End", mr.End, "OldBoardPosition", mr.OldBoardPosition)
	engineBoardPosition, found := gameLookup[mr.OldBoardPosition.UUID]
	if !found {
		return false, mr.OldBoardPosition, errors.New("cannot find game")
	}

	startRow, startCol, err := convertUICoordsToEngineCoords(mr.Start)
	if err != nil {
		fmt.Println("Problem converting UI coords")
		return false, mr.OldBoardPosition, err
	}
	endRow, endCol, _ := convertUICoordsToEngineCoords(mr.End)
	rowDy := endRow - startRow
	colDx := endCol - startCol

	p := engineBoardPosition.State[startRow][startCol]
	success, engineNewBoardPosition, _ := p.Move(rowDy, colDx, *engineBoardPosition)

	if !success {
		uiBoardPosition, _ := translateToUIBoardPosition(*engineBoardPosition)
		return false, uiBoardPosition, nil
	}

	engineNewBoardPosition.IsPlayer1Turn = !engineNewBoardPosition.IsPlayer1Turn
	gameLookup[engineNewBoardPosition.Uuid.String()] = &engineNewBoardPosition
	uiNewBoardPosition, _ := translateToUIBoardPosition(engineNewBoardPosition)

	return true, uiNewBoardPosition, nil
}

func main() {
	fmt.Println("Let's play chess!")

	_ = ui.StartUI(handleMove, handleFlip, handleInit)

}
