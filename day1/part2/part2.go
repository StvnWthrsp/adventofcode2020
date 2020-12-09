/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 1, Part 2

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

	fmt.Printf("Searching for integer triplet with sum = %s...\n", strconv.Itoa(sum))

	// In each iteration, check to see if the match to the current number is in the hash map
	// If found, print the pair and the multiplied value
	// If not found, add the current number to the hash map and continue
	for i := 0; i < len(dataArray)-2; i++ {
		numberI, err := strconv.Atoi(dataArray[i])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		tempI := 2020 - numberI

		for j := i+1; j < len(dataArray); j++ {
			numberJ, err := strconv.Atoi(dataArray[j])
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			tempJ := tempI - numberJ
			_, ok := m[tempJ]
			if ok {
				fmt.Printf("Found triplet: %d, %d and %d\n", numberI, numberJ, tempJ)
				answerInt := numberI * numberJ * tempJ
				fmt.Printf("Puzzle solution: %d*%d*%d = %d\n", numberI, numberJ, tempJ, answerInt)
				return
			}
			m[numberJ] = true
		}	
	}
}