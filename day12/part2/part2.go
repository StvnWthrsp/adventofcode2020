/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 12, Part 2

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
	"regexp"
	"math"
)

func main() {
	// Read input file
	inputData, err := ioutil.ReadFile(filepath.Join(os.Args[1]))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Split full file into array of individual lines
	dataArray := strings.Split(string(inputData), "\n")

	reDir := regexp.MustCompile(`^\w{1}`)
	reNum := regexp.MustCompile(`\d*$`)

	// Initialize the waypoint's initial position and the boat's initial position
	// We do not need to have the waypoint "move" with the boat, since the waypoint is always *relative* to the boat's position and the boat only moves to the waypoint
	// When the boat moves to the waypoint, we simply add the waypoint's x,y values to the boat's x,y value
	wx := 10
	wy := 1
	bx := 0
	by := 0

	for _, line := range dataArray {

		// If the instruction is a rotation, perform the appropriate Cartesian plane coordinate rotation
		if line == "L90" || line == "R270" {
			tempX := wx
			wx = -wy
			wy = tempX
			continue
		} else if line == "R90" || line == "L270" {
			tempX := wx
			wx = wy
			wy = -tempX
			continue
		} else if line == "R180" || line == "L180" {
			wx = -wx
			wy = -wy
			continue
		}

		num, _ := strconv.Atoi(reNum.FindString(line))
		dir := reDir.FindString(line)
		
		// All cardinal directions move only the waypoint. F moves only the boat.
		switch dir {
		case "N": wy += num
		case "E": wx += num
		case "S": wy -= num
		case "W": wx -= num
		case "F":
			bx += wx*num
			by += wy*num
		default: fmt.Printf("Unhandled direction: %s\n", line)
		}
	}
	fmt.Printf("Final Position: x=%d, y=%d\n", bx, by)
	fmt.Printf("Manhattan Position: |%d| + |%d| = %d\n", bx, by, int(math.Abs(float64(bx)) + math.Abs(float64(by))))
}