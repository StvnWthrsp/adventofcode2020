/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 9, Part 2

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

func getFirstInvalidNumber(dataArray []string) int {
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
				foundPair = true
				
				break
			}
		}

		// If a pair is not found, stop execution, and the last fmt.Printf gives the solution
		if !foundPair {
			return num
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
	return 0
}

func tryToSum(numToReach int, startingIndex int, dataArray []string) int {
	runningSum := 0
	for i := startingIndex; i < len(dataArray); i++ {
		line := dataArray[i]
		num, _ := strconv.Atoi(line)

		runningSum += num

		// If the sum we need is exceeded, call the function again with the next index
		if (runningSum) > numToReach {
			return tryToSum(numToReach, startingIndex+1, dataArray)
		}
		if (runningSum) == numToReach {
			fmt.Printf("Contiguous range with sum %d found.\n", numToReach)
			low := numToReach
			for j := startingIndex; j <= i; j++ {
				testNum, _ := strconv.Atoi(dataArray[j])
				if testNum < low {
					low = testNum
				}
			}
			hi := 0
			for j := startingIndex; j <= i; j++ {
				testNum, _ := strconv.Atoi(dataArray[j])
				if testNum > hi {
					hi = testNum
				}
			}
			solution := low + hi
			fmt.Printf("Sum of lowest and highest in contiguous range: %d + %d = %d\n", low, hi, solution)
			return solution
		}
	}
	return 0
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

	numToReach := getFirstInvalidNumber(dataArray)

	fmt.Printf("First Invalid Number: %d\n", numToReach)
	
	fmt.Printf("----- SOLUTION: %d -----\n", tryToSum(numToReach, 0, dataArray))
	
}