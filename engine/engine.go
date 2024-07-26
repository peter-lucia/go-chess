package engine

import (
	"github.com/google/uuid"
)

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

func NewBoard() (Board, error) {

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
	board.CheckMate = false
	board.Winner = "none"

	board.Uuid, _ = uuid.NewUUID()

	board.State[7][0] = Piece{CellType: P2Rook}
	board.State[7][1] = Piece{CellType: P2Horse}
	board.State[7][2] = Piece{CellType: P2Bishop}
	board.State[7][3] = Piece{CellType: P2Queen}
	board.State[7][4] = Piece{CellType: P2King}
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
