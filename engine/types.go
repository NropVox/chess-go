package engine

type Board [8][8]rune

type CastleAvailability struct {
	WhiteKing  bool
	WhiteQueen bool
	BlackKing  bool
	BlackQueen bool
}

type Coords struct {
	row int
	col int
}

type Chess struct {
	boardTable  Board
	turn        rune
	castle      CastleAvailability
	pawnPassant string
	halfmoves   int
	fullmoves   int

	winner rune
}
