/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 8, Part 1

This program takes the first argument as the input filepath
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	"regexp"
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
	
	acc := 0
	visitedLines := make(map[int]bool)

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		// If key i is present in the map, the instruction has already executed once. If key i is not present, add it to the map
		_, ok := visitedLines[i]
		if ok {
			break
		}
		visitedLines[i] = true

		// acc instruction: add the number to the acc variable and continue
		if strings.Contains(line, "acc") {
			re := regexp.MustCompile(`[\+|\-]\d*`)
			numStr := re.FindString(line)
			num, _ := strconv.Atoi(numStr)
			acc += num
			continue
		}

		// jmp instruction: add the number to i to move to the correct instruction (subtract one since the loop will also iterate it once)
		if strings.Contains(line, "jmp") {
			re := regexp.MustCompile(`[\+|\-]\d*`)
			numStr := re.FindString(line)
			num, _ := strconv.Atoi(numStr)
			i = i + num - 1
			continue
		}

		// No need to handle nop since it goes to the next line, which this loop already does
	}

	fmt.Printf("Value of acc before any instruction is executed a second time: %d\n", acc)

}