package main

import (
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
	"github.com/peter-lucia/go-chess/ui"
)

func handleMove(mr ui.MoveRequest) (bool, ui.BoardPosition, error) {
	// Handle chess move logic here
	fmt.Println("Start", mr.Start, "End", mr.End)
	// TODO: Return the entire board after the move request
	// with a move detail
	return false, mr.NewBoardPosition, nil
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
