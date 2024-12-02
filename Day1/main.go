package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
);
const fileLength int = 1000

func main() {
    file, err := os.Open("input")
    defer file.Close()

    r := bufio.NewReader(file)

	if err != nil {
        panic(err)
    }

    listA := [fileLength]int{}
    listB := [fileLength]int{}
    i := 0;

    for i < fileLength {
        line, _, err := r.ReadLine()
        if len(line) > 0 {
            tuple := strings.Fields(string(line))
            listA[i], _ = strconv.Atoi(tuple[0])
            listB[i], _ = strconv.Atoi(tuple[1])
            i++
        }
        if err != nil {
            break;
        }
    }
    sortedListA := bubbleSort(listA)
    sortedListB := bubbleSort(listB)

    //Part 1
    var sum int = sum(sortedListA, sortedListB)
    fmt.Println("Sum:", sum)

    //Part 2
    similarityScore := similarity(listA, listB)
    fmt.Println("Similarity score:", similarityScore)
}

func bubbleSort(arr [fileLength]int) [fileLength]int {
    length := len(arr)

    for i := 0; i < length; i++ {
        var swapped bool = false
        for j := 0; j < length; j++ {
            if(j+1 >= fileLength){
                continue
            }
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        if(!swapped) {
            break
        }
    }

    return arr
}

func sum(arrA [fileLength]int, arrB [fileLength]int) int {
    sum := 0;
    for i := 0; i< fileLength; i++ {
        if(arrA[i] > arrB[i]) {
            sum += (arrA[i] - arrB[i])
        } else {
            sum += (arrB[i] - arrA[i])
        }
    }
    return sum;
}

func similarity(arrA [fileLength]int, arrB [fileLength]int) int {
    similarityScore := 0
    for i := range(arrA) {
        similarityScore += countOccurances(arrA[i], arrB) * arrA[i]
    }
    return similarityScore
}

func countOccurances(number int, arr [fileLength]int) int {
    var sum int = 0;
    for i := range(arr) {
        if arr[i] == number {
            sum++;
        }
    }
    return sum
}