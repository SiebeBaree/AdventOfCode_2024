package main

import (
	"day5/utils"
	"fmt"
)

func getCenterElement(page []int) int {
	return page[len(page)/2]
}

func createRuleMap(rules [][]int) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		ruleMap[y] = append(ruleMap[y], x)
	}
	return ruleMap
}

func isValidOrder(pages []int, ruleMap map[int][]int) bool {
	positions := make(map[int]int)
	for i, page := range pages {
		positions[page] = i
	}

	for i, page := range pages {
		if rules, exists := ruleMap[page]; exists {
			for _, mustBeBefore := range rules {
				if pos, exists := positions[mustBeBefore]; exists && pos > i {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	rules, books, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	ruleMap := createRuleMap(rules)
	total := 0
	for _, pages := range books {
		if isValidOrder(pages, ruleMap) {
			total += getCenterElement(pages)
		}
	}

	fmt.Println(total)
}
