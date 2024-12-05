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

	var rows = []string{}
	count := 0

	for reader.Scan() {
		var line = reader.Text()
		rows = append(rows, line)
	}

	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i])-1; j++ {
			if rows[i][j] == 'A' {
				upperLeft := rows[i-1][j-1]
				lowerRight := rows[i+1][j+1]
				upperRight := rows[i-1][j+1]
				lowerLeft := rows[i+1][j-1]

				topLeftToBottomRight := string(upperLeft) + string(rows[i][j]) + string(lowerRight)
				bottomLeftToTopRight := string(lowerLeft) + string(rows[i][j]) + string(upperRight)

				if checkIfValidElements(topLeftToBottomRight, bottomLeftToTopRight) {
					count++
				}
			}
		}
	}

	fmt.Println("Results: ", count)
}

func checkIfValidElements(inputOne string, inputTwo string) bool {
	return (inputOne == "MAS" || inputOne == "SAM") && (inputTwo == "MAS" || inputTwo == "SAM")
}
