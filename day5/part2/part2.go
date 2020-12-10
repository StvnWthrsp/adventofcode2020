/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 4, Part 2

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
	
	minSeatId := 900
	maxSeatId := 0
	
	allSeatIds := make(map[int]bool)

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]
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
		seatId := rLow * 8 + cLow;
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
	for i := minSeatId; i < maxSeatId; i++ {
		_, ok := allSeatIds[i]
		if !ok {
			fmt.Printf("Missing Seat Id: %d\n", i)
		}
	}
}