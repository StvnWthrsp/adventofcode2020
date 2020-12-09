/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 1, Part 1

This program takes the first argument as the input filepath
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	sum := 2020
	m := make(map[int]bool)

	// Read input file
	inputData, err := ioutil.ReadFile(filepath.Join(os.Args[1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Split full file into array of individual lines (input has one integer per line)
	dataArray := strings.Split(string(inputData), "\n")

	fmt.Printf("Searching for integer pair with sum = %s...\n", strconv.Itoa(sum))

	// In each iteration, check to see if the match to the current number is in the hash map
	// If found, print the pair and the multiplied value
	// If not found, add the current number to the hash map and continue
	for i := 0; i < len(dataArray); i++ {
		number, err := strconv.Atoi(dataArray[i])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		temp := 2020 - number
		_, ok := m[temp]
		if ok {
			fmt.Printf("Found pair: %s and %s\n", dataArray[i], strconv.Itoa(temp))
			answerInt := number * temp
			fmt.Printf("Puzzle solution: %d*%d = %d\n", number, temp, answerInt)
			return
		}
		m[number] = true
	}
}