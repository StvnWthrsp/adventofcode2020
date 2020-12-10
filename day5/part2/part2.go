/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 5, Part 2

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
	
	// We need the minimum and maximum seat ID because seats are missing at the front and back. These values will tell us where to begin and end our search.
	minSeatId := 900
	maxSeatId := 0
	
	allSeatIds := make(map[int]bool)

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

		// Add the seat ID to the allSeatIds map and check whether it is lower than minimum / greater than maximum
		allSeatIds[int(seatId)] = true
		if int(seatId) > maxSeatId {
			maxSeatId = int(seatId)
		}
		if int(seatId) < minSeatId {
			minSeatId = int(seatId)
		}
	}

	fmt.Printf("Lowest Seat ID: %d\n", minSeatId)
	fmt.Printf("Highest Seat ID: %d\n", maxSeatId)

	// Search the map for keys between the minimum and maximum to find the one missing key, which is our seat ID
	for i := minSeatId; i < maxSeatId; i++ {
		_, ok := allSeatIds[i]
		if !ok {
			fmt.Printf("Missing Seat Id: %d\n", i)
		}
	}
}