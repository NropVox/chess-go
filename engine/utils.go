package engine

import "unicode"

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