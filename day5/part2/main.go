package main

import (
	"day5/utils"
	"fmt"
)

// Creates a simple map of dependencies: x -> y means x must come before y
func createRuleMap(rules [][]int) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		ruleMap[x] = append(ruleMap[x], y)
	}
	return ruleMap
}

func isValidOrder(pages []int, ruleMap map[int][]int) bool {
	positions := make(map[int]int)
	for i, page := range pages {
		positions[page] = i
	}

	for x, ys := range ruleMap {
		xPos, hasX := positions[x]
		if !hasX {
			continue
		}

		for _, y := range ys {
			yPos, hasY := positions[y]
			if hasY && xPos > yPos {
				return false
			}
		}
	}
	return true
}

func sortPages(pages []int, ruleMap map[int][]int) []int {
	result := make([]int, len(pages))
	copy(result, pages)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				// If we find a rule violation, swap the elements
				if mustComeBefore(result[j], result[i], ruleMap) {
					result[i], result[j] = result[j], result[i]
					changed = true
				}
			}
		}
	}
	return result
}

func mustComeBefore(x, y int, ruleMap map[int][]int) bool {
	if rules, exists := ruleMap[x]; exists {
		for _, rule := range rules {
			if rule == y {
				return true
			}
		}
	}
	return false
}

func main() {
	rules, books, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	ruleMap := createRuleMap(rules)
	sum := 0

	for _, pages := range books {
		if !isValidOrder(pages, ruleMap) {
			sorted := sortPages(pages, ruleMap)
			sum += sorted[len(sorted)/2]
		}
	}

	fmt.Println(sum)
}
