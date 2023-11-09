package main

import (
	"fmt"
	"github.com/peter-lucia/go-chess/engine"
)

func main() {
	fmt.Println("Let's play chess!")

	board, _ := engine.InitGame()

	for _, row := range board.State {
		for _, col := range row {
			fmt.Printf("%02d ", col.CellType) // %02d will pad 0's to make the width 2
		}
		fmt.Println()

	}

}
