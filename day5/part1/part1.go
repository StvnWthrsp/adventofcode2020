/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 5, Part 1

This program takes the first argument as the input filepath
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	"math"
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
	
	maxSeatId := 0.0

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]
		
		// Use first 7 characters to get seat row
		rLow := 0.0
		rHigh := 127.0
		for k := 0; k < 7; k++ {
			diff := rHigh - rLow
			switch string(line[k]) {
			case "F":
				rHigh = rLow + math.Floor(diff/2)
			case "B":
				rLow = rLow + math.Ceil(diff/2)
			}
		}

		// Use last 3 characters to get seat column
		cLow := 0.0
		cHigh := 7.0
		for k := 7; k < 10; k++ {
			diff := cHigh - cLow
			switch string(line[k]) {
			case "L":
				cHigh = cLow + math.Floor(diff/2)
			case "R":
				cLow = cLow + math.Ceil(diff/2)
			}
		}

		// Given expression to calculate seat ID
		seatId := rLow * 8 + cLow;

		// Keep track of highest seat ID for final answer
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	fmt.Printf("Highest Seat ID: %f\n", maxSeatId)
}