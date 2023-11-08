package main

import (
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
)

func main() {
	fmt.Println("Let's play chess!")

	board := engine.Board{}

	board.State[0][0] = engine.Piece{CellType: engine.P1Rook}
	board.State[0][1] = engine.Piece{CellType: engine.P1Horse}
	board.State[0][2] = engine.Piece{CellType: engine.P1Bishop}
	board.State[0][3] = engine.Piece{CellType: engine.P1Queen}
	board.State[0][4] = engine.Piece{CellType: engine.P1King}
	board.State[0][5] = engine.Piece{CellType: engine.P1Bishop}
	board.State[0][6] = engine.Piece{CellType: engine.P1Horse}
	board.State[0][7] = engine.Piece{CellType: engine.P1Rook}

	board.State[7][0] = engine.Piece{CellType: engine.P1Rook}
	board.State[7][1] = engine.Piece{CellType: engine.P1Horse}
	board.State[7][2] = engine.Piece{CellType: engine.P1Bishop}
	board.State[7][4] = engine.Piece{CellType: engine.P1King}
	board.State[7][3] = engine.Piece{CellType: engine.P1Queen}
	board.State[7][5] = engine.Piece{CellType: engine.P1Bishop}
	board.State[7][6] = engine.Piece{CellType: engine.P1Horse}
	board.State[7][7] = engine.Piece{CellType: engine.P1Rook}

	for row := range board.State {
		if row == 1 || row == 6 {
			for col := range board.State[row] {
				board.State[row][col] = engine.Piece{CellType: engine.P1Pawn}
			}
		}
	}

	for _, row := range board.State {
		for _, col := range row {
			fmt.Printf("%d ", col.CellType)
		}
		fmt.Println()

	}

}
