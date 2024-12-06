package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var fileArr [][]string
	var currentPosition [2]int
	var direction int

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		fileArr = append(fileArr, line)
		for col, char := range line {
			if char == "^" || char == "v" || char == "<" || char == ">" {
				currentPosition = [2]int{col, row}
				switch char {
				case "^":
					direction = 0
				case ">":
					direction = 1
				case "v":
					direction = 2
				case "<":
					direction = 3
				}
				fileArr[row][col] = "."
			}
		}
		row++
	}

	visited := make(map[[2]int]bool)
	visited[currentPosition] = true

	for {
		newPosition := move(currentPosition, direction)
		if !validPosition(newPosition, fileArr) {
			break
		}

		if fileArr[newPosition[1]][newPosition[0]] == "#" {
			direction = (direction + 1) % 4
		} else {
			currentPosition = newPosition
			visited[currentPosition] = true
		}
	}

	fmt.Println(len(visited))
}

func validPosition(pos [2]int, area [][]string) bool {
	return pos[1] >= 0 && pos[1] < len(area) && pos[0] >= 0 && pos[0] < len(area[0])
}

func move(pos [2]int, direction int) [2]int {
	switch direction {
	case 0:
		return [2]int{pos[0], pos[1] - 1}
	case 1:
		return [2]int{pos[0] + 1, pos[1]}
	case 2:
		return [2]int{pos[0], pos[1] + 1}
	case 3:
		return [2]int{pos[0] - 1, pos[1]}
	}
	return pos
}