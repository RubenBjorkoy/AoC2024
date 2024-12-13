package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	diskMap := []string{}
	individualBlocks := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diskMap = strings.Split(scanner.Text(), "")
	}

	// pointer := 0
	for i, val := range(diskMap) {
		currentValue, _ := strconv.Atoi(val)
		if i % 2 == 0 {
			//This means this ID contains a file
			for j := 0; j < currentValue; j++ { 
				individualBlocks = append(individualBlocks, strconv.Itoa(i/2))
			}
		} else if i % 2 == 1 {
			//This means this ID contians free-space
			for j := 0; j < currentValue; j++ { 
				individualBlocks = append(individualBlocks, ".")
			}
		}
	}

	for i := (len(individualBlocks) - 1); i > 0; i-- {
		if(individualBlocks[i] == ".") {continue}
		for j := 0; j < i; j++ {
			if(individualBlocks[j] != ".") {continue}
			individualBlocks[j] = individualBlocks[i]
			individualBlocks[i] = "."
			break
		}
	}

	sum := 0

	for i, val := range(individualBlocks) {
		if val == "." {break}
		num, _ := strconv.Atoi(val)
		sum += num * i 
	}

	fmt.Println(sum)
}