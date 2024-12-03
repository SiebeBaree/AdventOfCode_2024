package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var memory string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		memory += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return memory, nil
}
