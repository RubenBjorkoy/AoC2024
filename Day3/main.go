package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewScanner(file)

	sum1 := 0

	active := true 
	sum2 := 0

	for r.Scan() {
		var line = r.Text()
		re := regexp.MustCompile(`(do\(\)|don\'t\(\)|mul\([0-9]+,[0-9]+\))`)
		results := re.FindAllStringSubmatch(line, -1)
		for i := range(results) {
			if(results[i][0] == "do()") {
				active = true
			} else if(results[i][0] == "don't()") {
				active = false
			} else {
				sum1 += multiply(results[i][0])
				if(active) {
					sum2 += multiply(results[i][0])
				}
			}
		}
	}

	fmt.Println("Sum for task 1:", sum1)
	fmt.Println("Sum for task 2:", sum2)
}

func multiply(str string) int {
	re := regexp.MustCompile("[0-9]+")
	intArr :=  re.FindAllString(str, -1)
	int1, err1 := strconv.Atoi(intArr[0])
	int2, err2 := strconv.Atoi(intArr[1])
	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	}
	return int1 * int2
}