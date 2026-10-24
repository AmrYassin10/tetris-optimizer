package main

import (
	"fmt"
	"os"
	"strings"
)

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
	data := string(file)
	data = strings.ReplaceAll(data, "\r\n", "\n")
	data = strings.TrimSuffix(data, "\n")
	parts := strings.Split(data, "\n\n")

	var pieces [][]string

	for _, part := range parts {
		if !validateTetromino(part) {
			fmt.Println("ERROR")
			return
		}
		rows := strings.Split(part, "\n")
		pieces = append(pieces, rows)

	}
	size := 2
	for size*size < len(pieces)*4 {
		size++
	}
	for {

		board := make([][]rune, size)
		for i := 0; i < size; i++ {
			board[i] = make([]rune, size)
			for j := 0; j < size; j++ {
				board[i][j] = '.'
			}
		}
		if solve(board, pieces, 0) {
			for i := 0; i < size; i++ {
				fmt.Println(string(board[i]))
			}
			return
		}
		size++
	}
}
func validateTetromino(piece string) bool {
	rows := strings.Split(piece, "\n")
	if len(rows) != 4 {
		return false
	}
	for _, row := range rows {
		if len(row) != 4 {
			return false
		}
	}
	counter := 0
	for _, row := range rows {
		for _, char := range row {
			if char == '#' {
				counter++
			}
			if char != '#' && char != '.' {
				return false
			}

		}
	}
	if counter != 4 {
		return false
	}
	startRow := -1
	startCol := -1
	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
	visited := make([][]bool, 4)

	for i := range visited {
		visited[i] = make([]bool, 4)
	}
outer:
	for rowIndex, row := range rows {
		for colIndex, ch := range row {
			if ch == '#' {
				startRow = rowIndex
				startCol = colIndex
				break outer
			}
		}
	}
	queue := [][2]int{{startRow, startCol}}

	visited[startRow][startCol] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		row := current[0]
		col := current[1]
		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]
			if newRow < 0 || newRow > 3 || newCol < 0 || newCol > 3 {
				continue
			}
			if rows[newRow][newCol] == '#' && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}
	connected := 0
	for rowIndex, row := range rows {
		for colIndex, ch := range row {
			if ch == '#' && visited[rowIndex][colIndex] {

				connected++
			}
		}
	}
	return connected == 4

}
func canPlace(board [][]rune, piece []string, placementRow, placementCol int) bool {
	size := len(board)
	for rowIndex, row := range piece {
		for colIndex, ch := range row {
			if ch == '#' {
				boardRow := placementRow + rowIndex
				boardCol := placementCol + colIndex
				if boardRow < 0 || boardRow >= size ||
					boardCol < 0 || boardCol >= size {
					return false
				}
				if board[boardRow][boardCol] != '.' {
					return false
				}

			}
		}
	}
	return true
}
func placePiece(board [][]rune, piece []string, placementRow, placementCol int, label rune) {
	for rowIndex, row := range piece {
		for colIndex, ch := range row {
			if ch == '#' {
				boardRow := placementRow + rowIndex
				boardCol := placementCol + colIndex
				board[boardRow][boardCol] = label
			}
		}
	}
}
func removePiece(board [][]rune, piece []string, placementRow, placementCol int) {
	for rowIndex, row := range piece {
		for colIndex, ch := range row {
			if ch == '#' {
				boardRow := placementRow + rowIndex
				boardCol := placementCol + colIndex
				board[boardRow][boardCol] = '.'
			}
		}
	}
}
func solve(board [][]rune, pieces [][]string, pieceIndex int) bool {
	size := len(board)
	if pieceIndex == len(pieces) {
		return true
	}
	piece := pieces[pieceIndex]
	label := rune('A' + pieceIndex)

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if canPlace(board, piece, row, col) {
				placePiece(board, piece, row, col, label)
				if solve(board, pieces, pieceIndex+1) {
					return true
				}
				removePiece(board, piece, row, col)
			}
		}
	}

	return false
}
