/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 10, Part 1

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

func printChain(m map[int]bool, currentJolts int, totalOne *int, totalThree *int) (*int, *int) {
	nextJolts := currentJolts + 1

	_, ok := m[nextJolts]
	if ok {
		*totalOne = *totalOne + 1
		printChain(m, nextJolts, totalOne, totalThree)
	} else {
		nextJolts++
		_, ok := m[nextJolts]
		if ok {
			printChain(m, nextJolts, totalOne, totalThree)
		} else {
			nextJolts++
			_, ok := m[nextJolts]
			if ok {
				*totalThree = *totalThree + 1
				printChain(m, nextJolts, totalOne, totalThree)
			}
		}
	}
	return totalOne, totalThree
}

func main() {
	// Read input file
	inputData, err := ioutil.ReadFile(filepath.Join(os.Args[1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Split full file into array of individual lines (input has one integer per line)
	dataArray := strings.Split(string(inputData), "\n")

	m := make(map[int]bool)

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]
		num, _ := strconv.Atoi(line)

		m[num] = true
	}
	
	// Use pointers to update the running total values o and t as the function executes
	// Start t at 1 because there is always a 3-jolt gap between the last adapter and the built-in adapter
	o := 0
	t := 1
	one := &o
	three := &t
	printChain(m, 0, one, three)

	fmt.Printf("Gaps of 1: %d\nGaps of 3: %d\n", o, t)
	fmt.Printf("Solution: %d*%d = %d\n", o, t, o*t)
}