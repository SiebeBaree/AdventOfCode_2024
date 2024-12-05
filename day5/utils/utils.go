package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func ParseFile(filename string) ([][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var rules [][]int
	var books [][]int
	scanner := bufio.NewScanner(file)

	is_books := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			is_books = true
			continue
		}

		if is_books {
			books = append(books, stringsToInts(strings.Split(scanner.Text(), ",")))
		} else {
			rules = append(rules, stringsToInts(strings.Split(scanner.Text(), "|")))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return rules, books, nil
}
