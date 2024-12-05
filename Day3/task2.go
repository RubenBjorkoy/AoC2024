package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Task2() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewScanner(file)

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
				if(active) {
					sum2 += multiply(results[i][0])
				}
			}
		}
	}

	fmt.Println("Sum for task 2:", sum2)
}