package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Task1()
	Task2()
}

// Parse an update line into a slice of integers
func parseUpdate(line string) []int {
	parts := strings.Split(line, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			panic(fmt.Sprintf("Invalid page number: %s", part))
		}
		result = append(result, num)
	}
	return result
}

// Validate if an update follows the rules
func isUpdateValid(update []int, rules map[int][]int) bool {
	pageOrder := make(map[int]int)
	for i, page := range update {
		pageOrder[page] = i
	}

	for x, ys := range rules {
		for _, y := range ys {
			indexX, hasX := pageOrder[x]
			indexY, hasY := pageOrder[y]
			if hasX && hasY && indexX >= indexY {
				return false
			}
		}
	}

	return true
}

// Get the middle page of an update
func getMiddlePage(update []int) int {
	return update[len(update)/2]
}
