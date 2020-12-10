/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 6, Part 2

This program takes the first argument as the input filepath
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
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
	// Since the logic will calculate totals when it reaches a blank line, we must append an empty string to the end of the array
	dataArray = append(dataArray, "")

	answers := make(map[string]bool)
	total := 0
	
	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if len(line) == 0 {
			// If the line is blank, this is the end of the group, calculate the total number of answers (which is equal to the number of keys in the answers map)
			groupTotal := 0
			for k := range answers {
				fmt.Printf("%s", k)
				groupTotal++
				delete(answers, k)
			}
			total += groupTotal
			continue
		}

		if i == 0 {
			// At the start of the file, it is the beginning of a new group, so add each key to the map
			for j := 0; j < len(line); j++ {
				letter := string(line[j])
				answers[letter] = true
			}
			continue
		}

		if len(dataArray[i-1]) == 0 {
			// If the map is empty, it indicates the start of a new group, add each key to the map
			for j := 0; j < len(line); j++ {
				letter := string(line[j])
				answers[letter] = true
			}
			continue
		}

		// If the line is not blank, not the beginning of the file, and not the start of a new group, remove any keys NOT on the line
		for k := range answers {
			if !strings.Contains(line, k) {
				delete(answers, k)
			}
		}
	}
	
	fmt.Printf("Sum of questions answered by EVERYONE in a given group: %d\n", total)
}