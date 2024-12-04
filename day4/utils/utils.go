package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var puzzle [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		puzzle = append(puzzle, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return puzzle, nil
}
