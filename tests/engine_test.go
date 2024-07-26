package test_engine

import (
	"github.com/peter-lucia/go-chess/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckMate(t *testing.T) {
	board, err := engine.NewBoard()
	if err != nil {
		t.Error(err)
	}

	var moves [][]string
	moves = append(moves, []string{"f2", "f3"})
	moves = append(moves, []string{"e7", "e5"})
	moves = append(moves, []string{"g2", "g4"})
	moves = append(moves, []string{"d8", "h4"})

	for _, move := range moves {
		row1, col1, _ := engine.ConvertUICoordsToEngineCoords(move[0])
		row2, col2, _ := engine.ConvertUICoordsToEngineCoords(move[1])

		p := board.State[row1][col1]
		success, _, err := p.Move(row1, col1, row2-row1, col2-col1, &board)
		board.IsPlayer1Turn = !board.IsPlayer1Turn

		if err != nil {
			t.Error(err)
		}

		if success != true {
			t.Errorf("Failed to move %s to %s", move[0], move[1])
		}

	}

	assert.True(t, board.CheckMate)
	assert.Equal(t, "Player 1", board.Winner)

}
