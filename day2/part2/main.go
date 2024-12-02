package main

import (
	"day2/utils"
	"fmt"
)

func isSequenceValid(sequence []int) bool {
	if len(sequence) <= 1 {
		return true
	}

	is_increasing := sequence[0] < sequence[1]

	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]

		if is_increasing {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || -diff > 3 {
				return false
			}
		}
	}
	return true
}

func main() {
	reports, err := utils.ParseFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	total_safe_reports := 0
	for _, report := range reports {
		is_safe := isSequenceValid(report)

		if !is_safe {
			for i := 0; i < len(report); i++ {
				new_sequence := make([]int, 0)
				new_sequence = append(new_sequence, report[:i]...)
				new_sequence = append(new_sequence, report[i+1:]...)

				if isSequenceValid(new_sequence) {
					is_safe = true
					break
				}
			}
		}

		if is_safe {
			total_safe_reports++
		}
	}

	fmt.Println(total_safe_reports)
}
