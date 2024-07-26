package engine

type Cell int

const (
	Empty Cell = iota

	P1King
	P1Queen
	P1Bishop
	P1Horse
	P1Rook
	P1Pawn

	P2King
	P2Queen
	P2Bishop
	P2Horse
	P2Rook
	P2Pawn
	// be sure to update moveReachesEmptyCellOrOpponent before adding values here

)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)
