package engine

import (
	"unicode"
)

var (
	clearBoard = Board{
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-'},
	}
)

// translateCBtoCoords translates chessboard notation to coordinates
func translateCBtoCoords(cb string) *Coords {
	column := rune(cb[0])
	// Check if column is within range
	if column > 104 || column < 97 {
		return nil
	}

	row := rune(cb[1])
	// Check if row is within range
	if row > 56 || row < 49 {
		return nil
	}

	return &Coords{int('8' - row), int(column - 'a')}
}

// determineColor returns the color of the entered piece
func determineColor(piece rune) rune {
	if piece == '-' {
		return piece
	}

	if piece == unicode.ToUpper(piece) {
		return 'w'
	} else {
		return 'b'
	}
}

// determinePieceWithCoords determines the piece on the board with the coordinates
func determinePieceWithCoords(coord *Coords, board Board) rune {
	if checkIfCoordsIsOutOfBounds(coord) {
		return '-'
	}

	return board[coord.row][coord.col]
}

// checkIfCoordsIsOutOfBounds
func checkIfCoordsIsOutOfBounds(coord *Coords) bool {
	return coord.row > 7 || coord.row < 0 || coord.col > 7 || coord.col < 0
}

// checkIfThereIsPieceInCoords
func checkIfThereIsPieceInCoords(coord *Coords, board Board) bool {
	return determinePieceWithCoords(coord, board) != '-'
}

// checkIfAllyInCoords
func checkIfAllyInCoords(coord *Coords, color rune, board Board) bool {
	colorInCoord := determineColor(determinePieceWithCoords(coord, board))

	return colorInCoord == color
}

// determineColorPiece returns the enemy equivalent of piece
func determineColorPiece(color, piece rune) rune {
	if color == 'w' {
		return unicode.ToUpper(piece)
	} else {
		return unicode.ToLower(piece)
	}
}

// determineEnemyVersion returns the enemy version of the piece
func determineEnemyVersion(piece rune) rune {
	if unicode.IsUpper(piece) {
		return unicode.ToLower(piece)
	} else {
		return unicode.ToUpper(piece)
	}
}

// movePiece moves piece in board
func movePiece(from *Coords, to *Coords, board *Board) {
	board[to.row][to.col] = board[from.row][from.col]
	board[from.row][from.col] = '-'
}

// translateCoordsToCB translates coordinates to chessboard notation
func translateCoordsToCB(coord *Coords) string {
	if coord.row < 0 || coord.row > 7 || coord.col < 0 || coord.col > 7 {
		return ""
	}

	return string('a'+rune(coord.col)) + string('8'-rune(coord.row))
}

// checkIfMovesContains checks if moves contains move
func checkIfMovesContains(moves *[]*Coords, move *Coords) bool {
	for _, m := range *moves {
		if *m == *move {
			return true
		}
	}

	return false
}

// determineEnemy determines the enemy color
func determineEnemy(color rune) rune {
	if color == 'w' {
		return 'b'
	}
	return 'w'
}

// determineKingRow determines the row of the king
func determineKingRow(color rune) int {
	kingRow := 0
	if color == 'w' {
		kingRow = 7
	}
	return kingRow
}
