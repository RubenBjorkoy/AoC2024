package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
        panic(err)
    }


	defer file.Close()

	reader := bufio.NewScanner(file)

	safeCountTask1 := 0
	safeCountTask2 := 0
	for reader.Scan() {
		var strArr = strings.Fields(reader.Text())
		//Task 1
		safeCountTask1 += isSafe(strArr)

		//Task 2
		safeCountTask2 += isTolerablySafe(strArr)
	}

	fmt.Println("Amount of safe reports for task 1:", safeCountTask1)

	fmt.Println("Amount of safe reports for task 2:", safeCountTask2)
}

func isSafe(strArr []string) int {
	var reportSlice []int
	var prevVal int
	var increasing bool

	for i := range(strArr) {
		val, err := strconv.Atoi(strArr[i])
		if(err != nil) {panic(err)}
		reportSlice = append(reportSlice, val)

		if(i == 0) {
			prevVal = val
			nextVal, nextErr := strconv.Atoi(strArr[1])
			if(nextErr != nil) {panic(nextErr)}
			if(nextVal == val) {
				return 0
			}
			if(nextVal > val) {
				increasing = true
			} else {
				increasing = false
			}
			continue
		}
		if(increasing) {
			if(prevVal >= val || Abs(prevVal - val) > 3) {
				return 0
			}
		} else {
			if(prevVal <= val || Abs(prevVal - val) > 3) {
				return 0
			}
		}

		prevVal = val
	}

	return 1
}

func isTolerablySafe(strArr []string) int {
	isAlreadySafe := isSafe(strArr)
	if(isAlreadySafe == 1) {
		return 1
	}

	for i := range(strArr) {
		newStrArr := Remove(strArr, i)
		isSafeThisTime := isSafe(newStrArr)
		if(isSafeThisTime == 1) {
			return 1
		}
	}

	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Remove(slice []string, s int) []string {
	newSlice := make([]string, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}