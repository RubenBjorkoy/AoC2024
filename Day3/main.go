package main

import (
	"regexp"
	"strconv"
)

func main() {
	Task1()
	Task2()
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