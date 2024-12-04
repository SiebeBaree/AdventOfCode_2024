package main

import (
	"day4/utils"
	"fmt"
)

func isInBounds(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func checkDiagonalPattern(puzzle [][]string, centerRow, centerCol, dr, dc int) bool {
	rows := len(puzzle)
	cols := len(puzzle[0])

	mRow, mCol := centerRow-dr, centerCol-dc
	sRow, sCol := centerRow+dr, centerCol+dc

	if !isInBounds(mRow, mCol, rows, cols) || !isInBounds(sRow, sCol, rows, cols) {
		return false
	}

	return (puzzle[mRow][mCol] == "M" && puzzle[sRow][sCol] == "S") ||
		(puzzle[mRow][mCol] == "S" && puzzle[sRow][sCol] == "M")
}

func countXMASPatterns(puzzle [][]string, centerRow, centerCol int) bool {
	if puzzle[centerRow][centerCol] != "A" {
		return false
	}

	// Check main diagonal (top-left to bottom-right)
	mainDiagonal := checkDiagonalPattern(puzzle, centerRow, centerCol, 1, 1)

	// Check other diagonal (top-right to bottom-left)
	otherDiagonal := checkDiagonalPattern(puzzle, centerRow, centerCol, 1, -1)

	// If both diagonals have valid patterns, we found an X-MAS
	return mainDiagonal && otherDiagonal
}

func main() {
	puzzle, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	totalPatterns := 0
	for row := range puzzle {
		for col := range puzzle[row] {
			if countXMASPatterns(puzzle, row, col) {
				totalPatterns++
			}
		}
	}

	fmt.Println(totalPatterns)
}
