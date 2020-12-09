/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 2, Part 1

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

	count := 0

	for i := 0; i < len(dataArray); i++ {
		// Split string into distinct parts needed for comparison
		policy := strings.Split(dataArray[i], ": ")[0]
		password := strings.Split(dataArray[i], ": ")[1]
		occurences := strings.Split(policy, " ")[0]
		letter := strings.Split(policy, " ")[1]
		minOccurences, err := strconv.Atoi(strings.Split(occurences, "-")[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		maxOccurences, err := strconv.Atoi(strings.Split(occurences, "-")[1])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		
		// Check that letter occurs in password a number of times between minOccurences and maxOccurences
		if (strings.Count(password, letter) >= minOccurences) && strings.Count(password, letter) <= maxOccurences {
			count++
		}
	}
	fmt.Printf("Number of valid passwords: %d\n", count)
}