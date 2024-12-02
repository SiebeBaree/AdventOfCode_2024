package main

import (
	"day1/utils"
	"fmt"
)

func main() {
	left_list, right_list, err := utils.ParseFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	number := 0
	for _, left_value := range left_list {
		count := 0
		for _, right_value := range right_list {
			if left_value == right_value {
				count++
			}
		}

		number += left_value * count
	}

	fmt.Println(number)
}
