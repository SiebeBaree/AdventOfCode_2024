package main

import (
	"day1/utils"
	"fmt"
)

func Min(values []int) (value int, index int) {
	value = values[0]
	index = 0
	for i, v := range values {
		if v < value {
			value = v
			index = i
		}
	}

	return value, index
}

func main() {
	left_list, right_list, err := utils.ParseFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	total_iterations := len(left_list)
	distance := make([]int, len(left_list))

	for i := 0; i < total_iterations; i++ {
		smallest_number_left, index := Min(left_list)
		left_list = append(left_list[:index], left_list[index+1:]...)

		smallest_number_right, index := Min(right_list)
		right_list = append(right_list[:index], right_list[index+1:]...)

		if smallest_number_right > smallest_number_left {
			distance[i] = smallest_number_right - smallest_number_left
		} else {
			distance[i] = smallest_number_left - smallest_number_right
		}
	}

	total_distance := 0
	for _, value := range distance {
		total_distance += value
	}

	fmt.Println(total_distance)
}
