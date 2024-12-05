package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
        panic(err)
    }


	defer file.Close()

	reader := bufio.NewScanner(file)

	// var direction = 0
	crossWord := []string{}
	solution := [][][]int{} //A three dimensional array that contains the index of 'x' and 'm' from a complete 'xmas' solution location
	var currentLocation = [2]int{0, 0}
	const searchTerm = "XMAS"

	for reader.Scan() {
		var line = reader.Text()
		crossWord = append(crossWord, line)
	}
	
	for i := range(crossWord) {
		for j := range(crossWord[i]) {
			letter := crossWord[i][j]
			currentWord := [][2]int{}
			if(letter == 'X') {
				currentLocation = [2]int{j, i}
				fmt.Println("Found the letter ", string(letter))
				fmt.Println("at ", currentLocation[0], ",", currentLocation[1])
				currentWord = append(currentWord, currentLocation)
				var nextLetter [2]int
				
				for (len(currentWord) < len(searchTerm)) {
					nextLetter = findLetter(crossWord, currentLocation, searchTerm[len(currentWord)])
					if(nextLetter != [2]int{-1, -1}) {
						currentWord = append(currentWord, nextLetter)
					} else {
						break
					}
				}
			}
			if len(currentWord) == len(searchTerm) {
				convertedWord := make([][]int, len(currentWord))
				for k, v := range currentWord {
					convertedWord[k] = v[:]
				}
				solution = append(solution, convertedWord)
			}
		}
	}

	fmt.Println(len(solution))
}

func findLetter(crossWord []string, currentLocation [2]int, letter byte) [2]int {
	for y := -1; y<=1; y++ {
		if(currentLocation[1] + y < 0 || currentLocation[1] + y >= len(crossWord)) {continue}
		for x := -1; x<=1; x++ {
			if(currentLocation[0] + x < 0 || currentLocation[0] + x >= len(crossWord[currentLocation[1] + y])) {continue}
			if crossWord[currentLocation[1] + y][currentLocation[0] + x] == letter {
				fmt.Println("Found the letter ", string(letter))
				fmt.Println("at ", currentLocation[0] + x, ",", currentLocation[1] + y)
				return [2]int{currentLocation[0] + x, currentLocation[1] + y}
			}
		}
	}
	return [2]int{-1, -1}
}