/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 11, Part 2

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
func changeSeats(dataString string, lineLength int, lineNum int) (string, bool) {
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
		if c == 'L' && countFirstOccupied(x, y, lineLength, lineNum, byteData) == 0 {
			newData[i] = '#'
			didChange = true
		} else if c == '#' && countFirstOccupied(x, y, lineLength, lineNum, byteData) >= 5 {
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
	return string(newData), didChange
}

// Return true only if the given seat is "#"
func isOccupied(seat byte) bool {
	if seat == '#' {
		return true
	}
	return false
}

// Return true if the given seat is "L" or "#"
func isSeat(seat byte) bool {
	if seat == 'L' || seat == '#' {
		return true
	}
	return false
}

// Calculate a seat's raw index using the x and y position
func getIndex(lineLength int, lineNum int, x int, y int) int {
	return x + (y*lineLength)
}

// This is the main helper function of this program. It performs 8 while loops to search for the first seat in each direction, and returns the count of those seats which are occupied.
func countFirstOccupied(x int, y int, lineLength int, lineNum int, byteData []byte) int {
	count := 0
	searchY := y
	searchX := x

	// Search North
	for searchY > 0 {
		searchY--
		i := getIndex(lineLength, lineNum, x, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchY = y

	// Search South
	for searchY < lineNum-1 {
		searchY++
		i := getIndex(lineLength, lineNum, x, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchY = y

	// Search East
	for searchX < lineLength-1 {
		searchX++
		i := getIndex(lineLength, lineNum, searchX, y)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX = x

	// Search West
	for searchX > 0 {
		searchX--
		i := getIndex(lineLength, lineNum, searchX, y)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX = x

	// Search Northwest
	for searchX > 0 && searchY > 0 {
		searchX--
		searchY--
		i := getIndex(lineLength, lineNum, searchX, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX, searchY = x, y

	// Search Southeast
	for searchX < lineLength-1 && searchY < lineNum-1 {
		searchX++
		searchY++
		i := getIndex(lineLength, lineNum, searchX, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX, searchY = x, y

	// Search Southwest
	for searchX > 0 && searchY < lineNum-1 {
		searchX--
		searchY++
		i := getIndex(lineLength, lineNum, searchX, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX, searchY = x, y

	// Search Northeast
	for searchX < lineLength-1 && searchY > 0 {
		searchX++
		searchY--
		i := getIndex(lineLength, lineNum, searchX, searchY)
		if isOccupied(byteData[i]) {
			count++
			break
		} else if isSeat(byteData[i]) {
			break
		}
	}
	searchX, searchY = x, y
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
	nextDataString, changing := changeSeats(dataString, lineLength, lineNum)

	// Continue calling changeSeats until the map does not change
	for changing {
		nextDataString, changing = changeSeats(nextDataString, lineLength, lineNum)
	}

	// Count occupied seats identified by the # character
	re := regexp.MustCompile(`#`)
	totalCount := len(re.FindAllStringIndex(nextDataString, -1))
	fmt.Printf("Occupied Seats: %d\n", totalCount)
	
}