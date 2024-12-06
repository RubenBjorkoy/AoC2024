package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task1() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse the input rules
	rules := make(map[int][]int)
	var updates [][]int
	isReadingRules := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isReadingRules = false
			continue
		}

		if isReadingRules {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				panic(fmt.Sprintf("Invalid rule format: %s", line))
			}
			x, err1 := strconv.Atoi(parts[0])
			y, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic(fmt.Sprintf("Invalid rule numbers: %s", line))
			}
			rules[x] = append(rules[x], y)
		} else {
			update := parseUpdate(line)
			updates = append(updates, update)
		}
	}

	// Check updates and calculate the sum of middle pages
	sumOfMiddlePages := 0
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			middle := getMiddlePage(update)
			sumOfMiddlePages += middle
		}
	}

	fmt.Println("Task 1:", sumOfMiddlePages)
}
