package main

import (
	"day6/utils"
	"fmt"
)

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

	for !guardReachEdge(&content, guard_start_location) {
		moveGuard(&content, guard_start_location)
	}

	total_distance := 0
	for _, row := range content {
		for _, cell := range row {
			if cell == "X" {
				total_distance++
			}
		}
	}

	// +1 because the guard is not on the edge at the end
	fmt.Println(total_distance + 1)
}
