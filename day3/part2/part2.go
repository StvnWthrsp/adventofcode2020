/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 3, Part 2

This program takes the first argument as the input filepath
*/

/*
Notes about today's puzzle:
It isn't a matter of simply wrapping the input. Each LINE must "extend". This will require some modulo math, such as N % len(line)
If the length of a line is 7, then the index in the line is N % 7 where N is the horizontal position
Ex: N=3, N=6, N=9, N=12, N=15, ...
	i=3, i=6, i=2, i=5, i=1
Because we are only moving down 1 each time, we can simply iterate through an array of lines, checking if N % len(line) hits a tree
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)

// Given an array of strings (lines) and two integers corresponding to the slope (right, down), calculate the number of trees hit
func countTreesHit(dataArray []string, sr int, sd int) int {
	treeCount := 0
	rightPos := 0

	for downPos := sd; downPos < len(dataArray); downPos += sd {
		rightPos += sr
		linePos := rightPos % len(dataArray[downPos])
		if string(dataArray[downPos][linePos]) == "#" {
			treeCount++
		}
	}
	fmt.Printf("Number of trees hit for slope Right %d, Down %d: %d\n", sr, sd, treeCount)
	return treeCount
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

	totalCount := countTreesHit(dataArray, 1, 1) * countTreesHit(dataArray, 3, 1) * countTreesHit(dataArray, 5, 1) * countTreesHit(dataArray, 7, 1) * countTreesHit(dataArray, 1, 2)
	fmt.Printf("Total number of trees hit: %d\n", totalCount)
}