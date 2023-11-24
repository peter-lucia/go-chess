package main

import (
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
	"github.com/peter-lucia/go-chess/ui"
)

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

func convertPiece(uiPiece string, row int, col int, b *engine.Board) (engine.Board, error) {
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

	b := engine.Board{}
	b, _ = convertPiece(uiBoard.A8, 0, 0, &b)
	b, _ = convertPiece(uiBoard.B8, 0, 1, &b)
	b, _ = convertPiece(uiBoard.C8, 0, 2, &b)
	b, _ = convertPiece(uiBoard.D8, 0, 3, &b)
	b, _ = convertPiece(uiBoard.E8, 0, 4, &b)
	b, _ = convertPiece(uiBoard.F8, 0, 5, &b)
	b, _ = convertPiece(uiBoard.G8, 0, 6, &b)
	b, _ = convertPiece(uiBoard.H8, 0, 7, &b)

	b, _ = convertPiece(uiBoard.A7, 1, 0, &b)
	b, _ = convertPiece(uiBoard.B7, 1, 1, &b)
	b, _ = convertPiece(uiBoard.C7, 1, 2, &b)
	b, _ = convertPiece(uiBoard.D7, 1, 3, &b)
	b, _ = convertPiece(uiBoard.E7, 1, 4, &b)
	b, _ = convertPiece(uiBoard.F7, 1, 5, &b)
	b, _ = convertPiece(uiBoard.G7, 1, 6, &b)
	b, _ = convertPiece(uiBoard.H7, 1, 7, &b)

	b, _ = convertPiece(uiBoard.A6, 2, 0, &b)
	b, _ = convertPiece(uiBoard.B6, 2, 1, &b)
	b, _ = convertPiece(uiBoard.C6, 2, 2, &b)
	b, _ = convertPiece(uiBoard.D6, 2, 3, &b)
	b, _ = convertPiece(uiBoard.E6, 2, 4, &b)
	b, _ = convertPiece(uiBoard.F6, 2, 5, &b)
	b, _ = convertPiece(uiBoard.G6, 2, 6, &b)
	b, _ = convertPiece(uiBoard.H6, 2, 7, &b)

	b, _ = convertPiece(uiBoard.A5, 3, 0, &b)
	b, _ = convertPiece(uiBoard.B5, 3, 1, &b)
	b, _ = convertPiece(uiBoard.C5, 3, 2, &b)
	b, _ = convertPiece(uiBoard.D5, 3, 3, &b)
	b, _ = convertPiece(uiBoard.E5, 3, 4, &b)
	b, _ = convertPiece(uiBoard.F5, 3, 5, &b)
	b, _ = convertPiece(uiBoard.G5, 3, 6, &b)
	b, _ = convertPiece(uiBoard.H5, 3, 7, &b)

	b, _ = convertPiece(uiBoard.A4, 4, 0, &b)
	b, _ = convertPiece(uiBoard.B4, 4, 1, &b)
	b, _ = convertPiece(uiBoard.C4, 4, 2, &b)
	b, _ = convertPiece(uiBoard.D4, 4, 3, &b)
	b, _ = convertPiece(uiBoard.E4, 4, 4, &b)
	b, _ = convertPiece(uiBoard.F4, 4, 5, &b)
	b, _ = convertPiece(uiBoard.G4, 4, 6, &b)
	b, _ = convertPiece(uiBoard.H4, 4, 7, &b)

	b, _ = convertPiece(uiBoard.A3, 5, 0, &b)
	b, _ = convertPiece(uiBoard.B3, 5, 1, &b)
	b, _ = convertPiece(uiBoard.C3, 5, 2, &b)
	b, _ = convertPiece(uiBoard.D3, 5, 3, &b)
	b, _ = convertPiece(uiBoard.E3, 5, 4, &b)
	b, _ = convertPiece(uiBoard.F3, 5, 5, &b)
	b, _ = convertPiece(uiBoard.G3, 5, 6, &b)
	b, _ = convertPiece(uiBoard.H3, 5, 7, &b)

	b, _ = convertPiece(uiBoard.A2, 6, 0, &b)
	b, _ = convertPiece(uiBoard.B2, 6, 1, &b)
	b, _ = convertPiece(uiBoard.C2, 6, 2, &b)
	b, _ = convertPiece(uiBoard.D2, 6, 3, &b)
	b, _ = convertPiece(uiBoard.E2, 6, 4, &b)
	b, _ = convertPiece(uiBoard.F2, 6, 5, &b)
	b, _ = convertPiece(uiBoard.G2, 6, 6, &b)
	b, _ = convertPiece(uiBoard.H2, 6, 7, &b)

	b, _ = convertPiece(uiBoard.A1, 7, 0, &b)
	b, _ = convertPiece(uiBoard.B1, 7, 1, &b)
	b, _ = convertPiece(uiBoard.C1, 7, 2, &b)
	b, _ = convertPiece(uiBoard.D1, 7, 3, &b)
	b, _ = convertPiece(uiBoard.E1, 7, 4, &b)
	b, _ = convertPiece(uiBoard.F1, 7, 5, &b)
	b, _ = convertPiece(uiBoard.G1, 7, 6, &b)
	b, _ = convertPiece(uiBoard.H1, 7, 7, &b)

	return b, nil

}

func translateToUIBoardPosition(engineBoardPosition engine.Board) (ui.BoardPosition, error) {

	uiBoard := ui.BoardPosition{}
	uiBoard.A8, _ = convertToUIPiece(engineBoardPosition.State[0][0])
	uiBoard.B8, _ = convertToUIPiece(engineBoardPosition.State[0][1])
	uiBoard.C8, _ = convertToUIPiece(engineBoardPosition.State[0][2])
	uiBoard.D8, _ = convertToUIPiece(engineBoardPosition.State[0][3])
	uiBoard.E8, _ = convertToUIPiece(engineBoardPosition.State[0][4])
	uiBoard.F8, _ = convertToUIPiece(engineBoardPosition.State[0][5])
	uiBoard.G8, _ = convertToUIPiece(engineBoardPosition.State[0][6])
	uiBoard.H8, _ = convertToUIPiece(engineBoardPosition.State[0][7])

	uiBoard.A7, _ = convertToUIPiece(engineBoardPosition.State[1][0])
	uiBoard.B7, _ = convertToUIPiece(engineBoardPosition.State[1][1])
	uiBoard.C7, _ = convertToUIPiece(engineBoardPosition.State[1][2])
	uiBoard.D7, _ = convertToUIPiece(engineBoardPosition.State[1][3])
	uiBoard.E7, _ = convertToUIPiece(engineBoardPosition.State[1][4])
	uiBoard.F7, _ = convertToUIPiece(engineBoardPosition.State[1][5])
	uiBoard.G7, _ = convertToUIPiece(engineBoardPosition.State[1][6])
	uiBoard.H7, _ = convertToUIPiece(engineBoardPosition.State[1][7])

	uiBoard.A6, _ = convertToUIPiece(engineBoardPosition.State[2][0])
	uiBoard.B6, _ = convertToUIPiece(engineBoardPosition.State[2][1])
	uiBoard.C6, _ = convertToUIPiece(engineBoardPosition.State[2][2])
	uiBoard.D6, _ = convertToUIPiece(engineBoardPosition.State[2][3])
	uiBoard.E6, _ = convertToUIPiece(engineBoardPosition.State[2][4])
	uiBoard.F6, _ = convertToUIPiece(engineBoardPosition.State[2][5])
	uiBoard.G6, _ = convertToUIPiece(engineBoardPosition.State[2][6])
	uiBoard.H6, _ = convertToUIPiece(engineBoardPosition.State[2][7])

	uiBoard.A5, _ = convertToUIPiece(engineBoardPosition.State[3][0])
	uiBoard.B5, _ = convertToUIPiece(engineBoardPosition.State[3][1])
	uiBoard.C5, _ = convertToUIPiece(engineBoardPosition.State[3][2])
	uiBoard.D5, _ = convertToUIPiece(engineBoardPosition.State[3][3])
	uiBoard.E5, _ = convertToUIPiece(engineBoardPosition.State[3][4])
	uiBoard.F5, _ = convertToUIPiece(engineBoardPosition.State[3][5])
	uiBoard.G5, _ = convertToUIPiece(engineBoardPosition.State[3][6])
	uiBoard.H5, _ = convertToUIPiece(engineBoardPosition.State[3][7])

	uiBoard.A4, _ = convertToUIPiece(engineBoardPosition.State[4][0])
	uiBoard.B4, _ = convertToUIPiece(engineBoardPosition.State[4][1])
	uiBoard.C4, _ = convertToUIPiece(engineBoardPosition.State[4][2])
	uiBoard.D4, _ = convertToUIPiece(engineBoardPosition.State[4][3])
	uiBoard.E4, _ = convertToUIPiece(engineBoardPosition.State[4][4])
	uiBoard.F4, _ = convertToUIPiece(engineBoardPosition.State[4][5])
	uiBoard.G4, _ = convertToUIPiece(engineBoardPosition.State[4][6])
	uiBoard.H4, _ = convertToUIPiece(engineBoardPosition.State[4][7])

	uiBoard.A3, _ = convertToUIPiece(engineBoardPosition.State[5][0])
	uiBoard.B3, _ = convertToUIPiece(engineBoardPosition.State[5][1])
	uiBoard.C3, _ = convertToUIPiece(engineBoardPosition.State[5][2])
	uiBoard.D3, _ = convertToUIPiece(engineBoardPosition.State[5][3])
	uiBoard.E3, _ = convertToUIPiece(engineBoardPosition.State[5][4])
	uiBoard.F3, _ = convertToUIPiece(engineBoardPosition.State[5][5])
	uiBoard.G3, _ = convertToUIPiece(engineBoardPosition.State[5][6])
	uiBoard.H3, _ = convertToUIPiece(engineBoardPosition.State[5][7])

	uiBoard.A2, _ = convertToUIPiece(engineBoardPosition.State[6][0])
	uiBoard.B2, _ = convertToUIPiece(engineBoardPosition.State[6][1])
	uiBoard.C2, _ = convertToUIPiece(engineBoardPosition.State[6][2])
	uiBoard.D2, _ = convertToUIPiece(engineBoardPosition.State[6][3])
	uiBoard.E2, _ = convertToUIPiece(engineBoardPosition.State[6][4])
	uiBoard.F2, _ = convertToUIPiece(engineBoardPosition.State[6][5])
	uiBoard.G2, _ = convertToUIPiece(engineBoardPosition.State[6][6])
	uiBoard.H2, _ = convertToUIPiece(engineBoardPosition.State[6][7])

	uiBoard.A1, _ = convertToUIPiece(engineBoardPosition.State[7][0])
	uiBoard.B1, _ = convertToUIPiece(engineBoardPosition.State[7][1])
	uiBoard.C1, _ = convertToUIPiece(engineBoardPosition.State[7][2])
	uiBoard.D1, _ = convertToUIPiece(engineBoardPosition.State[7][3])
	uiBoard.E1, _ = convertToUIPiece(engineBoardPosition.State[7][4])
	uiBoard.F1, _ = convertToUIPiece(engineBoardPosition.State[7][5])
	uiBoard.G1, _ = convertToUIPiece(engineBoardPosition.State[7][6])
	uiBoard.H1, _ = convertToUIPiece(engineBoardPosition.State[7][7])

	return uiBoard, nil

}

func handleMove(mr ui.MoveRequest) (bool, ui.BoardPosition, error) {
	// Returns true, new board detail, nil if the move was successful
	// returns false, new board detail, nil if the move was a failure
	// returns an error if there was a problem with the move
	fmt.Println("Start", mr.Start, "End", mr.End, "NewBoardPosition", mr.NewBoardPosition)
	engineNewBoardPosition, _ := translateToEngineBoardPosition(mr.NewBoardPosition)

	// TODO: Logic on the new board position
	// ...

	uiNewBoardPosition, _ := translateToUIBoardPosition(engineNewBoardPosition)
	return true, uiNewBoardPosition, nil
}

func main() {
	fmt.Println("Let's play chess!")

	board, _ := engine.InitGame()

	fmt.Println("        Player 1")
	for _, row := range board.State {
		for _, col := range row {
			fmt.Printf("%02d ", col.CellType) // %02d will pad 0's to make the width 2
		}
		fmt.Println()

	}
	fmt.Println("        Player 2")
	ui.StartUI(handleMove)

}
