package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var leftColumn []int
	var rightColumn []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", scanner.Text())
		}

		left, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting left number: %v", err)
		}
		leftColumn = append(leftColumn, left)

		right, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting right number: %v", err)
		}
		rightColumn = append(rightColumn, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return leftColumn, rightColumn, nil
}
