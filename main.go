package main

import "fmt"

type Piece int

const (
	Queen Piece = iota
	King
	Bishop
	Horse
	Rook
	Pawn
)

func main() {
	fmt.Println("Let's play chess!")

	board := [8][8]Piece{}

	board[0][0] = Rook
	board[0][1] = Horse
	board[0][2] = Bishop
	board[0][3] = Queen
	board[0][4] = King
	board[0][5] = Bishop
	board[0][6] = Horse
	board[0][7] = Rook

	board[7][0] = Rook
	board[7][1] = Horse
	board[7][2] = Bishop
	board[7][4] = King
	board[7][3] = Queen
	board[7][5] = Bishop
	board[7][6] = Horse
	board[7][7] = Rook

	for i := range board[1] {
		board[1][i] = Pawn
	}
	for i := range board[6] {
		board[6][i] = Pawn
	}

	for _, row := range board {
		fmt.Println(row)
	}

}
