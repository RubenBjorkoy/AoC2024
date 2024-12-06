package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("testInput")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileArr [][]string
	direction := 3

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fileArr = append(fileArr, []string{line})
	}
	fmt.Println(direction)
	direction = changeDirection(direction)
	fmt.Println(direction)

	fmt.Println(fileArr)
}

func changeDirection(direction int) int {
	direction = (direction + 1) % 4
	return direction
}
