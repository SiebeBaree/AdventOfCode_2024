package main

import (
	"day3/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func extractMultiplications(input string) []string {
	pattern := `(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`
	re := regexp.MustCompile(pattern)
	return re.FindAllString(input, -1)
}

func parseMultiplication(multiplication string) (int, int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(multiplication)

	first, err1 := strconv.Atoi(matches[1])
	second, err2 := strconv.Atoi(matches[2])
	if err1 != nil || err2 != nil {
		return 0, 0
	}

	return first, second
}

func main() {
	memory, err := utils.ParseFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	valid_multiplications := extractMultiplications(memory)

	is_enabled := true
	sum := 0
	for _, multiplication := range valid_multiplications {

		if strings.HasPrefix(multiplication, "don't") {
			is_enabled = false
		} else if strings.HasPrefix(multiplication, "do") {
			is_enabled = true
		} else if is_enabled {
			first, second := parseMultiplication(multiplication)
			sum += first * second
		}
	}

	fmt.Println(sum)
}
