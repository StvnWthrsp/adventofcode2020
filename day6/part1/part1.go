/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 6, Part 1

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
				groupTotal++
				delete(answers, k)
			}
			total += groupTotal
			continue
		}
		// If the line is not blank, add each letter on the line as a key in the map
		for j := 0; j < len(line); j++ {
			letter := string(line[j])
			_, ok := answers[letter]
			if !ok {
				answers[letter] = true
			}
		}
	}

	fmt.Printf("Sum of questions answered by ANYONE in a given group: %d\n", total)
}