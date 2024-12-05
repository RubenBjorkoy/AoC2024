package main

import (
	"bufio"
	"fmt"
	"os"
)

func Task2() {
	file, err := os.Open("input")

	if err != nil {
        panic(err)
    }


	defer file.Close()

	reader := bufio.NewScanner(file)

	crossWord := []string{}
	solutions := [][][]int{}
	const searchTerm = "XMAS"

	for reader.Scan() {
		var line = reader.Text()
		crossWord = append(crossWord, line)
	}
	
	for i := range(crossWord) {
		for j := range(crossWord[i]) {
			letter := crossWord[i][j]
			if letter == 'X' {
				for _, initialDirection := range getDirections() {
					currentWord := [][2]int{}
					currentLocation := [2]int{j, i}
					currentWord = append(currentWord, currentLocation)

					for len(currentWord) < len(searchTerm) {
						nextLetter, newDirection := findLetter(crossWord, currentLocation, searchTerm[len(currentWord)], initialDirection)

						if nextLetter != [2]int{-1, -1} {
							currentWord = append(currentWord, nextLetter)
							currentLocation = nextLetter
							initialDirection = newDirection
						} else {
							break
						}
					}

					if len(currentWord) == len(searchTerm) {
						convertedWord := make([][]int, len(currentWord))
						for k, v := range currentWord {
							convertedWord[k] = v[:]
						}
						solutions = append(solutions, convertedWord)
					}
				}
			}
		}
	}

	fmt.Println("Total solutions found:", len(solutions))
}