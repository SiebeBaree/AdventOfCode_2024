package main

import (
	"day6/utils"
	"fmt"
)

func copyGrid(content [][]string) [][]string {
	newGrid := make([][]string, len(content))
	for i := range content {
		newGrid[i] = make([]string, len(content[i]))
		copy(newGrid[i], content[i])
	}
	return newGrid
}

func moveGuard(content *[][]string, guard_location []int) {
	row, col := guard_location[0], guard_location[1]
	current_direction := (*content)[row][col]

	var next_row, next_col int
	switch current_direction {
	case "^":
		next_row, next_col = row-1, col
	case "v":
		next_row, next_col = row+1, col
	case "<":
		next_row, next_col = row, col-1
	case ">":
		next_row, next_col = row, col+1
	}

	if next_row < 0 || next_row >= len(*content) || next_col < 0 || next_col >= len((*content)[0]) || (*content)[next_row][next_col] == "#" {
		switch current_direction {
		case "^":
			(*content)[row][col] = ">"
		case ">":
			(*content)[row][col] = "v"
		case "v":
			(*content)[row][col] = "<"
		case "<":
			(*content)[row][col] = "^"
		}
	} else {
		(*content)[next_row][next_col] = current_direction
		(*content)[row][col] = "X"
		guard_location[0], guard_location[1] = next_row, next_col
	}
}

func guardReachEdge(content *[][]string, guard_location []int) bool {
	row, col := guard_location[0], guard_location[1]
	current_direction := (*content)[row][col]

	switch current_direction {
	case "^":
		return row == 0
	case "v":
		return row == len(*content)-1
	case "<":
		return col == 0
	case ">":
		return col == len((*content)[0])-1
	}

	return false
}

func detectLoop(visited map[string]bool, content *[][]string, guard_location []int) bool {
	state := fmt.Sprintf("%d,%d,%s", guard_location[0], guard_location[1], (*content)[guard_location[0]][guard_location[1]])
	if visited[state] {
		return true
	}
	visited[state] = true
	return false
}

func main() {
	content, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	guard_start_location := []int{}
	for i, row := range content {
		for j, cell := range row {
			if cell == "^" {
				guard_start_location = append(guard_start_location, i, j)
			}
		}
	}

	loopPositions := 0
	for i := 0; i < len(content); i++ {
		for j := 0; j < len(content[0]); j++ {
			// Skip if position is already occupied or is guard's start position
			if content[i][j] != "." || (i == guard_start_location[0] && j == guard_start_location[1]) {
				continue
			}

			// Try placing obstruction at this position
			testGrid := copyGrid(content)
			testGrid[i][j] = "#"
			testGuardLoc := []int{guard_start_location[0], guard_start_location[1]}
			visited := make(map[string]bool)

			// Simulate guard movement until either reaching edge or detecting loop
			for !guardReachEdge(&testGrid, testGuardLoc) {
				if detectLoop(visited, &testGrid, testGuardLoc) {
					loopPositions++
					break
				}
				moveGuard(&testGrid, testGuardLoc)
			}
		}
	}

	fmt.Println(loopPositions)
}
