package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task2() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			rules[x] = append(rules[x], y)
		} else {
			update := parseUpdate(line)
			updates = append(updates, update)
		}
	}

	sumOfMiddlePages := 0
	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			sortedUpdate := sortUpdate(update, rules)
			middle := getMiddlePage(sortedUpdate)
			sumOfMiddlePages += middle
		}
	}

	fmt.Println("Task 2:", sumOfMiddlePages)
}

func sortUpdate(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	inUpdate := make(map[int]bool)

	for _, page := range update {
		inUpdate[page] = true
		inDegree[page] = 0
	}

	for x, ys := range rules {
		if inUpdate[x] {
			for _, y := range ys {
				if inUpdate[y] {
					graph[x] = append(graph[x], y)
					inDegree[y]++
				}
			}
		}
	}

	queue := []int{}
	for page := range inUpdate {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	fmt.Println(graph)

	return result
}