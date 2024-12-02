package main

import (
	"day2/utils"
	"fmt"
)

func main() {
	reports, err := utils.ParseFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	total_safe_reports := 0
	for _, report := range reports {
		is_safe := true
		is_increasing := report[1] > report[0]

		for i := 1; i < len(report); i += 1 {
			previous := report[i-1]
			current := report[i]

			if is_increasing && (previous >= current || current-previous > 3) {
				is_safe = false
				break
			} else if !is_increasing && (previous <= current || previous-current > 3) {
				is_safe = false
				break
			}
		}

		if is_safe {
			total_safe_reports++
		}
	}

	fmt.Println(total_safe_reports)
}
