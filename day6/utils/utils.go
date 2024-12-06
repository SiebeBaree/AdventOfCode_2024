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

	var content [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return content, nil
}
