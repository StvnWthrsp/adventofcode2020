/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 12, Part 1

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
	x := 0
	y := 0
	rot := 90

	for _, line := range dataArray {
		num, _ := strconv.Atoi(reNum.FindString(line))
		dir := reDir.FindString(line)
		// The modulo does not return the results I expected for negative numbers
		// A simple solution is adding 360 to negative results
		switch dir {
		case "N": y += num
		case "E": x += num
		case "S": y -= num
		case "W": x -= num
		case "R": rot += num
		case "L": rot -= num
		case "F":
			mod := rot % 360
			if mod < 0 {
				mod += 360
			}
			switch mod {
			case 0: y += num
			case 90: x += num
			case 180: y -= num
			case 270: x -= num
			}
		}
	}
	fmt.Printf("Final Position: x=%d, y=%d\n", x, y)
	fmt.Printf("Manhattan Position: |%d| + |%d| = %d\n", x, y, int(math.Abs(float64(x)) + math.Abs(float64(y))))
}