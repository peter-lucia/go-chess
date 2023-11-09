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
			fmt.Printf("%d ", col.CellType)
		}
		fmt.Println()

	}

}
