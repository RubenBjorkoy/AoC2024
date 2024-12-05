package main

import (
	"fmt"
)

func main() {
	Task1()
	Task2()
}

func findLetter(crossWord []string, currentLocation [2]int, letter byte, direction [2]int) ([2]int, [2]int) {
	newY := currentLocation[1] + direction[1]
	newX := currentLocation[0] + direction[0] 
	
	if newY < 0 || newY >= len(crossWord) || newX < 0 || newX >= len(crossWord[newY]) {
		return [2]int{-1, -1}, [2]int{-2, -2}
	}

	if crossWord[newY][newX] == letter {
		fmt.Printf("Found the letter '%c' at %d, %d\n", letter, newX, newY)
		return [2]int{newX, newY}, direction
	}

	fmt.Printf("Failed to find letter '%c' from location %d, %d in direction (%d, %d)\n", letter, currentLocation[0], currentLocation[1], direction[0], direction[1])
	return [2]int{-1, -1}, [2]int{-2, -2}
}

func getDirections() [][2]int {
	return [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0},           {1, 0},
		{-1, 1},  {0, 1},  {1, 1},
	}
}