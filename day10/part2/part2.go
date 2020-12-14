/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 10, Part 2

This one was tough! It's more a math puzzle, as any solution that would calculate EVERY combination would be extremely inefficient. 
The tribonacci sequence is key to solving this. The number of combinations for a given adapter is the sum of the number of combinations that can reach the previous 3 adapters. Why? Because each of the previous 3 adapters can jump up to the given adapter!
By writing out small samples on paper I also concluded that we need to replace missing adapters with a 0.

We can reach a 7-jolt adapter through every combination that will reach 4-, 5-, or 6-jolts. If we are missing the 5-jolt adapter, we can only reach 7-jolts through the combinations that can generate 4- or 6-jolts

This can also be viewed as dynamic programming. The solution to the full problem is the number of combinations that reach the maximum jolts adapter, and the solution to each subproblem is the number of combinations that reach the given adapter.

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

func getMaxJolts(m map[int]bool) int {
	maxJolts := 0
	for k := range m {
		if k > maxJolts {
			maxJolts = k
		}
	}
	fmt.Printf("Max Jolts: %d\n", maxJolts)
	return maxJolts
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

	maxJolts := getMaxJolts(m)
	
	// Map each adapter (k) to the running total of combinations (v) that can reach the adapter k
	allPaths := make(map[int]int)
	allPaths[0] = 1

	for i := 1; i <= maxJolts; i++ {
		_ = i
		// If we have a 1-jolt adapter, there will be 2 combinations that can reach the next adapter. Set allPaths[1] to 1 so that allPaths[0] + allPaths[1] will equal 2.
		if i == 1 {
			_, ok := m[i]
			if ok {
				allPaths[i] = 1
			}
		// This is just to keep the array index in bounds...
		} else if i == 2 {
			_, ok := m[i]
			if ok {
				allPaths[i] = allPaths[i-2] + allPaths[i-1]
			} else {
				allPaths[i] = 0
			}
		// Generate tribonacci sequence, but set allPaths[i] to 0 if the i-jolts adapter is missing
		} else {
			_, ok := m[i]
			if ok {
				allPaths[i] = allPaths[i-3] + allPaths[i-2] + allPaths[i-1]
			} else {
				allPaths[i] = 0
			}
		}
	}

	fmt.Printf("Total number of possible combinations: %d\n", allPaths[maxJolts])
}