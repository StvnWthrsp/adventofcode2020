/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 11, Part 1

This file currently prints out each iteration of the seap map, which is not necessary, but I enjoy the fun visual representation.

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
	"regexp"
)

// Iterates through every seat in dataString to apply the appropriate rules based on the number of occupied seats returned by countFirstOccupied. Returns the new seat data as a single string, and boolean indicating whether any seat changed.
func changeSeats(dataString string, lineLength int, lineNum int) ([]string, string, bool) {
	byteData := make([]byte, len(dataString))
	copy(byteData, dataString)
	newData := make([]byte, len(dataString))
	copy(newData, dataString)

	seatMap := make([]string, lineNum)
	didChange := false

	// Iterate through every seat
	for i, c := range byteData {
		// Calculate x and y grid positions to simplify math within the helper function
		x, y := i % lineLength, int(math.Floor(float64(i/lineLength)))

		// Count seats and apply rules
		if c == 'L' && countOccupiedAdjacent(x, y, lineLength, lineNum, byteData) == 0 {
			newData[i] = '#'
			didChange = true
		} else if c == '#' && countOccupiedAdjacent(x, y, lineLength, lineNum, byteData) >= 4 {
			newData[i] = 'L'
			didChange = true
		}

		// If we're at the end of a line, print it
		if x == lineLength-1 {
			fmt.Printf("%s\n", string(newData[i-lineLength+1:i+1]))
			seatMap[y] = string(newData[i-lineLength+1:i+1])
		}
	}
	fmt.Printf("---------------------\n")
	return seatMap, string(newData), didChange
}

// Return true only if the given seat is "#"
func isOccupied(seat byte) bool {
	if seat == '#' {
		return true
	}
	return false
}

// This is the main helper function of this program. This uses switch statements based on the x,y grid position to determine which seats need to be examined. Returns the count of seats adjacent to the seat at x,y.
func countOccupiedAdjacent(x int, y int, maxX int, maxY int, byteData []byte) int {
	index := x + (y*maxX)
	count := 0
	switch y {
	case 0:
		switch x {
		case 0:
			adj := []int{1, maxX, maxX+1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		case maxX-1:
			adj := []int{-1, maxX-1, maxX}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		default:
			adj := []int{-1, 1, maxX-1, maxX, maxX+1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		}
	case maxY-1:
		switch x {
		case 0:
			adj := []int{1, -maxX, -maxX+1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		case maxX-1:
			adj := []int{-1, -maxX, -maxX-1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		default:
			adj := []int{-1, 1, -maxX+1, -maxX, -maxX-1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		}
	default:
		switch x {
		case 0:
			adj := []int{1, -maxX, -maxX+1, maxX, maxX+1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		case maxX-1:
			adj := []int{-1, maxX, maxX-1, -maxX, -maxX-1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		default:
			adj := []int{-1, 1, maxX-1, maxX, maxX+1, -maxX+1, -maxX, -maxX-1}
			for _, v := range adj {
				if isOccupied(byteData[index+v]) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// Read input file
	inputData, err := ioutil.ReadFile(filepath.Join(os.Args[1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Split full file into array of individual lines to calculate lineLength and lineNum, then print the original seat map
	dataArray := strings.Split(string(inputData), "\n")
	lineLength := len(dataArray[0])
	lineNum := len(dataArray)

	for line := range dataArray {
		fmt.Printf("%s\n", dataArray[line])
	}
	fmt.Printf("---------------------\n")

	// The program will end when the seat map is no longer changing
	changing := true

	// We need a raw array of the data, not split into separate lines
	dataString := strings.Replace(string(inputData), "\n", "", -1)

	// First call to changeSeats, passing in the original data
	_, nextDataString, changing := changeSeats(dataString, lineLength, lineNum)

	// Continue calling changeSeats until the map does not change
	for changing {
		_, nextDataString, changing = changeSeats(nextDataString, lineLength, lineNum)
	}

	// Count occupied seats identified by the # character
	re := regexp.MustCompile(`#`)
	totalCount := len(re.FindAllStringIndex(nextDataString, -1))
	fmt.Printf("Occupied Seats: %d\n", totalCount)
	
}