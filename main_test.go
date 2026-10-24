package main

import "testing"

func TestValidateTetromino(t *testing.T) {
	piece := "....\n.##.\n..##\n...."

	got := validateTetromino(piece)

	if !got {
		t.Errorf("validateTetromino() = %v, want true", got)
	}

}
func TestValidateDisconnectedTetromino(t *testing.T) {
	piece := "#...\n....\n..##\n...#"

	got := validateTetromino(piece)

	if got {
		t.Errorf("validateTetromino() = %v, want false", got)
	}
}
func TestValidateWrongCellCount(t *testing.T) {
	piece := "....\n.##.\n...#\n...."

	got := validateTetromino(piece)

	if got {
		t.Errorf("validateTetromino() = %v, want false", got)
	}

}
func TestValidateInvalidCharacter(t *testing.T) {
	piece := "X...\n##..\n.#..\n...."
	got := validateTetromino(piece)

	if got {
		t.Errorf("validateTetromino() = %v, want false", got)
	}

}

func TestValidateWrongRowWidth(t *testing.T) {
	piece := "#..\n##..\n.#..\n...."
	got := validateTetromino(piece)

	if got {
		t.Errorf("validateTetromino() = %v, want false", got)
	}
}
func TestCanPlaceOnEmptyBoard(t *testing.T) {
	board := make([][]rune, 4)
	for i := 0; i < 4; i++ {
		board[i] = make([]rune, 4)
		for j := 0; j < 4; j++ {
			board[i][j] = '.'
		}
	}
	piece := []string{
		"##..",
		"##..",
		"....",
		"....",
	}
	got := canPlace(board, piece, 0, 0)
	if !got {
		t.Errorf("canPlace() = %v, want true", got)
	}
}
func TestCanPlaceOnOccupiedCells(t *testing.T) {
	board := make([][]rune, 4)
	for i := 0; i < 4; i++ {
		board[i] = make([]rune, 4)
		for j := 0; j < 4; j++ {
			board[i][j] = '.'
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			board[i][j] = 'A'
		}
	}
	piece := []string{
		"##..",
		"##..",
		"....",
		"....",
	}

	got := canPlace(board, piece, 0, 0)
	if got {
		t.Errorf("canPlace() = %v, want false", got)
	}

}
func TestCanPlaceOutsideBoard(t *testing.T) {
	board := make([][]rune, 4)
	for i := 0; i < 4; i++ {
		board[i] = make([]rune, 4)
		for j := 0; j < 4; j++ {
			board[i][j] = '.'
		}
	}
	piece := []string{
		"##..",
		"##..",
		"....",
		"....",
	}
	got := canPlace(board, piece, 3, 3)
	if got {
		t.Errorf("canPlace() = %v, want false", got)
	}

}
func TestSolveSingleTetromino(t *testing.T) {
	board := make([][]rune, 2)
	for i := 0; i < 2; i++ {
		board[i] = make([]rune, 2)
		for j := 0; j < 2; j++ {
			board[i][j] = '.'
		}
	}
	pieces := [][]string{
		{
			"##..",
			"##..",
			"....",
			"....",
		},
	}
	got := solve(board, pieces, 0)
	if !got {
		t.Errorf("solve()=%v,want true", got)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {

			if board[i][j] != 'A' {
				t.Errorf("board[%d][%d] = %c, want A",
					i,
					j,
					board[i][j])

			}

		}
	}
}

// Creates an empty square board filled with dots.
func createEmptyBoard(size int) [][]rune {
	board := make([][]rune, size)

	for i := 0; i < size; i++ {
		board[i] = make([]rune, size)

		for j := 0; j < size; j++ {
			board[i][j] = '.'
		}
	}

	return board
}

// Tests a tetromino with fewer than four rows.
func TestValidateWrongRowCount(t *testing.T) {
	piece := "....\n.##.\n..##"

	if validateTetromino(piece) {
		t.Error("validateTetromino() = true, want false for three rows")
	}
}

// Tests a tetromino containing more than four occupied cells.
func TestValidateTooManyCells(t *testing.T) {
	piece := "##..\n##..\n#...\n...."

	if validateTetromino(piece) {
		t.Error("validateTetromino() = true, want false for five occupied cells")
	}
}

// Tests a valid L-shaped tetromino.
func TestValidateLTetromino(t *testing.T) {
	piece := "#...\n#...\n##..\n...."

	if !validateTetromino(piece) {
		t.Error("validateTetromino() = false, want true for valid L tetromino")
	}
}

// Tests a valid square tetromino.
func TestValidateSquareTetromino(t *testing.T) {
	piece := "##..\n##..\n....\n...."

	if !validateTetromino(piece) {
		t.Error("validateTetromino() = false, want true for square tetromino")
	}
}

// Tests placing a piece on the board.
func TestPlacePiece(t *testing.T) {
	board := createEmptyBoard(4)
	piece := []string{
		"##..",
		"##..",
		"....",
		"....",
	}

	placePiece(board, piece, 0, 0, 'A')

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if board[i][j] != 'A' {
				t.Errorf("board[%d][%d] = %c, want A", i, j, board[i][j])
			}
		}
	}
}

// Tests removing a piece from the board.
func TestRemovePiece(t *testing.T) {
	board := createEmptyBoard(4)
	piece := []string{
		"##..",
		"##..",
		"....",
		"....",
	}

	placePiece(board, piece, 0, 0, 'A')
	removePiece(board, piece, 0, 0)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if board[i][j] != '.' {
				t.Errorf("board[%d][%d] = %c, want .", i, j, board[i][j])
			}
		}
	}
}

// Tests solving with two square tetrominoes.
func TestSolveTwoTetrominoes(t *testing.T) {
	board := createEmptyBoard(4)

	pieces := [][]string{
		{
			"##..",
			"##..",
			"....",
			"....",
		},
		{
			"##..",
			"##..",
			"....",
			"....",
		},
	}

	if !solve(board, pieces, 0) {
		t.Fatal("solve() = false, want true for two tetrominoes")
	}

	labels := make(map[rune]int)
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				labels[board[i][j]]++
			}
		}
	}

	if labels['A'] != 4 {
		t.Errorf("piece A occupies %d cells, want 4", labels['A'])
	}
	if labels['B'] != 4 {
		t.Errorf("piece B occupies %d cells, want 4", labels['B'])
	}
}

// Tests that the solver can assemble two straight tetrominoes
// in a 5x5 board, where a 4x4 board cannot fit them.
func TestSolveMinimumBoardForTwoStraightPieces(t *testing.T) {
	pieces := [][]string{
		{
			"#...",
			"#...",
			"#...",
			"#...",
		},
		{
			"####",
			"....",
			"....",
			"....",
		},
	}

	board := createEmptyBoard(5)

	if !solve(board, pieces, 0) {
		t.Fatal("solve() = false, want true on a 5x5 board")
	}

	labels := make(map[rune]int)
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				labels[board[i][j]]++
			}
		}
	}

	if labels['A'] != 4 {
		t.Errorf("piece A occupies %d cells, want 4", labels['A'])
	}
	if labels['B'] != 4 {
		t.Errorf("piece B occupies %d cells, want 4", labels['B'])
	}
}
