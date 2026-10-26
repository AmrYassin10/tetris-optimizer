package main

import (
	"strings"
	"testing"
)

func createTestBoard(size int) [][]rune {
	board := make([][]rune, size)
	for row := range board {
		board[row] = make([]rune, size)
		for col := range board[row] {
			board[row][col] = '.'
		}
	}
	return board
}

func TestValidateTetrominoValidSquare(t *testing.T) {
	input := "....\n.##.\n.##.\n...."

	piece, valid := validateTetromino(input)

	if !valid {
		t.Fatal("expected a valid tetromino")
	}

	if len(piece.Cells) != 4 {
		t.Fatalf("expected 4 occupied cells, got %d", len(piece.Cells))
	}

	if piece.Height != 2 || piece.Width != 2 {
		t.Fatalf(
			"expected trimmed dimensions 2x2, got %dx%d",
			piece.Height,
			piece.Width,
		)
	}

	expected := []Cell{
		{Row: 0, Col: 0},
		{Row: 0, Col: 1},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
	}

	for i, cell := range expected {
		if piece.Cells[i] != cell {
			t.Errorf("cell %d: expected %+v, got %+v", i, cell, piece.Cells[i])
		}
	}
}

func TestValidateTetrominoVerticalLine(t *testing.T) {
	input := "#...\n#...\n#...\n#..."

	piece, valid := validateTetromino(input)

	if !valid {
		t.Fatal("expected a valid vertical line")
	}

	if piece.Height != 4 || piece.Width != 1 {
		t.Fatalf(
			"expected dimensions 4x1, got %dx%d",
			piece.Height,
			piece.Width,
		)
	}
}

func TestValidateTetrominoWrongRowCount(t *testing.T) {
	input := "....\n.##.\n.##."

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected invalid tetromino with only 3 rows")
	}
}

func TestValidateTetrominoWrongRowWidth(t *testing.T) {
	input := "....\n.##.\n.##.\n..."

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected invalid tetromino with a row of incorrect width")
	}
}

func TestValidateTetrominoWrongCellCount(t *testing.T) {
	input := "....\n.##.\n.#..\n...."

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected invalid tetromino with only 3 occupied cells")
	}
}

func TestValidateTetrominoInvalidCharacter(t *testing.T) {
	input := "....\n.##.\n.##.\n...X"

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected invalid character to be rejected")
	}
}

func TestValidateTetrominoDisconnected(t *testing.T) {
	input := "#..#\n....\n....\n...."

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected disconnected cells to be rejected")
	}
}

func TestValidateTetrominoLShape(t *testing.T) {
	input := "#...\n#...\n##..\n...."

	piece, valid := validateTetromino(input)

	if !valid {
		t.Fatal("expected valid L-shaped tetromino")
	}

	if len(piece.Cells) != 4 {
		t.Fatalf("expected 4 cells, got %d", len(piece.Cells))
	}

	if piece.Height != 3 || piece.Width != 2 {
		t.Fatalf(
			"expected dimensions 3x2, got %dx%d",
			piece.Height,
			piece.Width,
		)
	}
}

func TestCanPlaceOnEmptyBoard(t *testing.T) {
	piece, valid := validateTetromino("....\n.##.\n.##.\n....")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(4)

	if !canPlace(board, piece, 0, 0) {
		t.Fatal("expected square to fit on empty board")
	}
}

func TestCanPlaceOnOccupiedCells(t *testing.T) {
	piece, valid := validateTetromino("....\n.##.\n.##.\n....")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(4)
	board[0][0] = 'A'

	if canPlace(board, piece, 0, 0) {
		t.Fatal("expected placement to fail due to occupied cell")
	}
}

func TestCanPlaceOutsideBoard(t *testing.T) {
	piece, valid := validateTetromino("#...\n#...\n#...\n#...")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(3)

	if canPlace(board, piece, 0, 0) {
		t.Fatal("expected vertical piece not to fit on a 3x3 board")
	}
}

func TestPlacePiece(t *testing.T) {
	piece, valid := validateTetromino("....\n.##.\n.##.\n....")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(4)

	placePiece(board, piece, 1, 1, 'A')

	expected := []Cell{
		{Row: 1, Col: 1},
		{Row: 1, Col: 2},
		{Row: 2, Col: 1},
		{Row: 2, Col: 2},
	}

	for _, cell := range expected {
		if board[cell.Row][cell.Col] != 'A' {
			t.Errorf(
				"expected A at (%d,%d)",
				cell.Row,
				cell.Col,
			)
		}
	}
}

func TestRemovePiece(t *testing.T) {
	piece, valid := validateTetromino("....\n.##.\n.##.\n....")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(4)
	placePiece(board, piece, 1, 1, 'A')
	removePiece(board, piece, 1, 1)

	for _, cell := range piece.Cells {
		if board[1+cell.Row][1+cell.Col] != '.' {
			t.Errorf(
				"expected empty cell at (%d,%d) after removal",
				1+cell.Row,
				1+cell.Col,
			)
		}
	}
}

func TestSolveSingleTetromino(t *testing.T) {
	piece, valid := validateTetromino("#...\n#...\n#...\n#...")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(4)

	if !solve(board, []Piece{piece}, 0) {
		t.Fatal("expected solver to place one tetromino")
	}

	count := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 'A' {
				count++
			}
		}
	}

	if count != 4 {
		t.Fatalf("expected 4 cells labelled A, got %d", count)
	}
}

func TestSolveTwoTetrominoes(t *testing.T) {
	pieceA, valid := validateTetromino("#...\n#...\n#...\n#...")
	if !valid {
		t.Fatal("first test piece should be valid")
	}

	pieceB, valid := validateTetromino("....\n.##.\n.##.\n....")
	if !valid {
		t.Fatal("second test piece should be valid")
	}

	board := createTestBoard(4)

	if !solve(board, []Piece{pieceA, pieceB}, 0) {
		t.Fatal("expected solver to place both tetrominoes")
	}

	countA, countB := 0, 0

	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case 'A':
				countA++
			case 'B':
				countB++
			}
		}
	}

	if countA != 4 {
		t.Errorf("expected 4 cells labelled A, got %d", countA)
	}

	if countB != 4 {
		t.Errorf("expected 4 cells labelled B, got %d", countB)
	}
}

func TestSolveCannotFit(t *testing.T) {
	piece, valid := validateTetromino("#...\n#...\n#...\n#...")
	if !valid {
		t.Fatal("test piece should be valid")
	}

	board := createTestBoard(3)

	if solve(board, []Piece{piece}, 0) {
		t.Fatal("expected solver to fail on a 3x3 board")
	}
}

func TestValidateTetrominoRejectsExtraInput(t *testing.T) {
	input := strings.Repeat("#", 5)

	_, valid := validateTetromino(input)

	if valid {
		t.Fatal("expected malformed input to be rejected")
	}
}

