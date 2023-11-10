package main

import (
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
	"github.com/peter-lucia/go-chess/ui"
)

//func handleMove(some parameters) {
//	// Handle chess move logic here
//}

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
	ui.StartUI()

}
