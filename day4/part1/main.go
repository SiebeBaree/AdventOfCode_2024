package main

import (
	"day4/utils"
	"fmt"
)

type Direction struct {
	rowDelta    int
	columnDelta int
}

func getAllDirections() []Direction {
	return []Direction{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // diagonal down-right
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
		{-1, 1},  // diagonal up-right
	}
}

func isInBounds(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

// Counts how many times "MAS" appears after X in all directions
func countXMASPatterns(puzzle [][]string, startRow, startCol int) int {
	patternCount := 0
	rows := len(puzzle)
	cols := len(puzzle[0])

	// Check each direction for the "MAS" pattern
	for _, dir := range getAllDirections() {
		// Calculate positions for M, A, and S
		mRow, mCol := startRow+dir.rowDelta, startCol+dir.columnDelta
		aRow, aCol := startRow+2*dir.rowDelta, startCol+2*dir.columnDelta
		sRow, sCol := startRow+3*dir.rowDelta, startCol+3*dir.columnDelta

		// Check if all positions are within bounds
		if !isInBounds(mRow, mCol, rows, cols) ||
			!isInBounds(aRow, aCol, rows, cols) ||
			!isInBounds(sRow, sCol, rows, cols) {
			continue
		}

		// Check if we found "MAS" pattern
		if puzzle[mRow][mCol] == "M" &&
			puzzle[aRow][aCol] == "A" &&
			puzzle[sRow][sCol] == "S" {
			patternCount++
		}
	}

	return patternCount
}

func main() {
	puzzle, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	totalPatterns := 0
	for rowIndex := range puzzle {
		for colIndex := range puzzle[rowIndex] {
			if puzzle[rowIndex][colIndex] == "X" {
				totalPatterns += countXMASPatterns(puzzle, rowIndex, colIndex)
			}
		}
	}

	fmt.Println(totalPatterns)
}
