package main

import (
	"bufio"
	"fmt"
	"math"
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

	scanner := bufio.NewScanner(file)
	var correctEquations [][]int

	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		var values []int
		for _, x := range(line) {
			x = replaceSemiColon(x)
			intX, _ := strconv.Atoi(x)
			values = append(values, intX)
		}
		if(validEquation(values)) {
			correctEquations = append(correctEquations, values)
		}
	}
	var sum int
	for _, x := range correctEquations {
		sum += x[0]
	}
	fmt.Println(sum)
}

func validEquation(equation []int) bool {
	solution := equation[0]
	var values []int
	for i := range equation {
		if(i>0) {
			values = append(values, equation[i])
		}
	}
	length := len(values)
	var permutations int = 2*powInt(2, length-2)

	for i := 0; i<permutations; i++ {
		operators := generateOperators(i, length-1)
		if evaluateEquation(values, operators) == solution {
			return true
		}
	}
	
	return false
}

func generateOperators(index int, length int) []rune {
	operators := make([]rune, length)
	for i := 0; i < length; i++ {
		if (index >> i)&1 == 0 {
			operators[i] = '+'
		} else {
			operators[i] = '*'
		}
	}
	return operators
}

func evaluateEquation(values []int, operators []rune) int {
	result := values[0]
	for i := 0; i<len(operators); i++ {
		if operators[i] == '+' {
			result += values[i+1]
		} else if operators[i] == '*' {
			result *= values[i+1]
		}
	}
	return result
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func replaceSemiColon(str string) string {
	for _, x := range str {
		if(x == ':') {
			return strings.Replace(str, ":", "", -1)
		}
	}
	return str
}