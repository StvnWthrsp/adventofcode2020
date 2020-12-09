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
		positions := strings.Split(policy, " ")[0]
		letter := strings.Split(policy, " ")[1]
		p1, err := strconv.Atoi(strings.Split(positions, "-")[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		p2, err := strconv.Atoi(strings.Split(positions, "-")[1])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// The puzzle does not start arrays at 0, so decerement p1 and p2
		p1--
		p2--
		
		// Either p1 OR p2 must be equal to letter (p1 XOR p2)
		if string(password[p1]) == letter && string(password[p2]) != letter {
			count++
		} else if string(password[p2]) == letter && string(password[p1]) != letter {
			count++
		}
	}
	fmt.Printf("Number of valid passwords: %d\n", count)
}