package main

import (
	"fmt"
	"os"
	"strings"
)

type Cell struct {
	Row int
	Col int
}

type Piece struct {
	Cells  []Cell
	Height int
	Width  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	data := strings.ReplaceAll(string(file), "\r\n", "\n")
	data = strings.TrimSuffix(data, "\n")

	parts := strings.Split(data, "\n\n")
	pieces := make([]Piece, 0, len(parts))

	for _, part := range parts {
		piece, valid := validateTetromino(part)
		if !valid {
			fmt.Println("ERROR")
			return
		}
		pieces = append(pieces, piece)
	}

	size := 2
	for size*size < len(pieces)*4 {
		size++
	}

	for {
		board := make([][]rune, size)
		for row := range board {
			board[row] = make([]rune, size)
			for col := range board[row] {
				board[row][col] = '.'
			}
		}

		if solve(board, pieces, 0) {
			for _, row := range board {
				fmt.Println(string(row))
			}
			return
		}

		size++
	}
}

func validateTetromino(input string) (Piece, bool) {
	var empty Piece

	rows := strings.Split(input, "\n")
	if len(rows) != 4 {
		return empty, false
	}

	for _, row := range rows {
		if len(row) != 4 {
			return empty, false
		}
	}

	var cells []Cell
	minRow, minCol := 4, 4
	maxRow, maxCol := -1, -1

	for rowIndex, row := range rows {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			ch := row[colIndex]

			if ch != '#' && ch != '.' {
				return empty, false
			}

			if ch == '#' {
				cells = append(cells, Cell{
					Row: rowIndex,
					Col: colIndex,
				})

				if rowIndex < minRow {
					minRow = rowIndex
				}
				if rowIndex > maxRow {
					maxRow = rowIndex
				}
				if colIndex < minCol {
					minCol = colIndex
				}
				if colIndex > maxCol {
					maxCol = colIndex
				}
			}
		}
	}

	if len(cells) != 4 {
		return empty, false
	}

	if !isConnected(rows, cells[0]) {
		return empty, false
	}

	// Remove empty borders by shifting occupied coordinates.
	for i := range cells {
		cells[i].Row -= minRow
		cells[i].Col -= minCol
	}

	return Piece{
		Cells:  cells,
		Height: maxRow - minRow + 1,
		Width:  maxCol - minCol + 1,
	}, true
}

func isConnected(rows []string, start Cell) bool {
	var visited [4][4]bool
	queue := []Cell{start}
	visited[start.Row][start.Col] = true

	directions := [4]Cell{
		{Row: -1, Col: 0},
		{Row: 1, Col: 0},
		{Row: 0, Col: -1},
		{Row: 0, Col: 1},
	}

	connected := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		connected++

		for _, direction := range directions {
			newRow := current.Row + direction.Row
			newCol := current.Col + direction.Col

			if newRow < 0 || newRow >= 4 ||
				newCol < 0 || newCol >= 4 {
				continue
			}

			if rows[newRow][newCol] != '#' ||
				visited[newRow][newCol] {
				continue
			}

			visited[newRow][newCol] = true
			queue = append(queue, Cell{
				Row: newRow,
				Col: newCol,
			})
		}
	}

	return connected == 4
}

func canPlace(
	board [][]rune,
	piece Piece,
	placementRow, placementCol int,
) bool {
	size := len(board)

	if placementRow+piece.Height > size ||
		placementCol+piece.Width > size {
		return false
	}

	for _, cell := range piece.Cells {
		boardRow := placementRow + cell.Row
		boardCol := placementCol + cell.Col

		if board[boardRow][boardCol] != '.' {
			return false
		}
	}

	return true
}

func placePiece(
	board [][]rune,
	piece Piece,
	placementRow, placementCol int,
	label rune,
) {
	for _, cell := range piece.Cells {
		board[placementRow+cell.Row][placementCol+cell.Col] = label
	}
}

func removePiece(
	board [][]rune,
	piece Piece,
	placementRow, placementCol int,
) {
	for _, cell := range piece.Cells {
		board[placementRow+cell.Row][placementCol+cell.Col] = '.'
	}
}

func solve(board [][]rune, pieces []Piece, pieceIndex int) bool {
	if pieceIndex == len(pieces) {
		return true
	}

	size := len(board)
	piece := pieces[pieceIndex]
	label := rune('A' + pieceIndex)

	for row := 0; row <= size-piece.Height; row++ {
		for col := 0; col <= size-piece.Width; col++ {
			if !canPlace(board, piece, row, col) {
				continue
			}

			placePiece(board, piece, row, col, label)

			if solve(board, pieces, pieceIndex+1) {
				return true
			}

			removePiece(board, piece, row, col)
		}
	}

	return false
}
