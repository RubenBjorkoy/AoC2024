package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Task1() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewScanner(file)

	sum1 := 0

	for r.Scan() {
		var line = r.Text()
		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		results := re.FindAllStringSubmatch(line, -1)
		for i := range(results) {
			sum1 += multiply(results[i][0])
		}
	}

	fmt.Println("Sum for task 1:", sum1)
}