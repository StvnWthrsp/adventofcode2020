/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 9, Part 1

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
	// Read input file
	inputData, err := ioutil.ReadFile(filepath.Join(os.Args[1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Split full file into array of individual lines (input has one integer per line)
	dataArray := strings.Split(string(inputData), "\n")

	m := make(map[int]bool)
	var prev []int

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]
		num, _ := strconv.Atoi(line)

		// For the first 25 lines just store the values in prev and m
		if i < 25 {
			prev = append(prev, num)
			m[num] = true
			continue
		}

		fmt.Printf("A pair from the previous 25 numbers must sum to: %d\n", num)
		foundPair := false

		// Iterate through the 25 previous numbers
		for j := 0; j < 25; j++ {	
			temp := num - prev[j]

			// Numbers must be unique - skip if the two numbers temp and prev[j] are equal
			if temp == prev[j] {
				continue
			}

			// Determine if the matching number temp is in map m containing the past 25 numbers
			_, ok := m[temp]
			if ok {
				fmt.Printf("Found pair: %d and %s\n", prev[j], strconv.Itoa(temp))
				fmt.Printf("-----------------------------------\n")
				foundPair = true
				
				break
			}
		}

		// If a pair is not found, stop execution, and the last fmt.Printf gives the solution
		if !foundPair {
			return
		}

		// Remove the oldest key from m and add the current num as a key
		delete(m, prev[0])
		m[num] = true

		// Remove the oldest number from prev by copying the all values except prev[0] into a new array, add num to the end of the new array, and copy it back into prev
		var newPrev [25]int
		copy(newPrev[:], prev[1:])
		newPrev[24] = num
		copy(prev[:], newPrev[:])
		
	}
}