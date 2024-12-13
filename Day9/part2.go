package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
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

	individualBlocks = mapOutBlocks(diskMap)
	individualBlocks = moveBlocksToLeftmostFreeSpace(individualBlocks)

	sum := 0

	for i, val := range(individualBlocks) {
		if val == "." {continue}
		num, _ := strconv.Atoi(val)
		sum += num * i 
	}

	fmt.Println(sum)
}

func mapOutBlocks(arr []string) []string {
	individualBlocks := []string{}
	for i, val := range(arr) {
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

	return individualBlocks
}

func moveBlocksToLeftmostFreeSpace(individualBlocks []string) []string {
	for i := (len(individualBlocks) - 1); i > 0; i-- {
		if(individualBlocks[i] == ".") {continue}
		blockSize := determineBlockSize(individualBlocks, i)
		for j := 0; j < i; j++ {
			if(individualBlocks[j] != ".") {continue}
			dotSize := determineBlockSize(individualBlocks, j)
			if(dotSize < blockSize) {
				j += dotSize
			} else {
				for x := 0; x < blockSize; x++ {
					individualBlocks[j + x] = individualBlocks[i - x]
					individualBlocks[i - x] = "."
				}
				i -= blockSize
				break
			}
		}
	}

	return individualBlocks
}

func determineBlockSize(arr []string, index int) int {
	size := 1

	for i := index + 1; i < len(arr); i++ {
		if i > len(arr) {break}
		if(arr[i] == arr[index]) {
			size++
		} else {
			break
		}
	}
	for i := index - 1; i >= 0; i-- {
		if i < 0 {break}
		if(arr[i] == arr[index]) {
			size++
		} else {
			break
		}
	}
	return size
}