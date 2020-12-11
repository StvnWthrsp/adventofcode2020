/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 8, Part 2

This program takes the first argument as the input filepath

LINE CHANGED IN INPUT.TXT
442 jmp +122
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
	
	visitedLines := make(map[int]bool)
	saveLoop := false
	var loop []int

	// First iteration through the instructions to determine the looping pattern, store it in slice loop
	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		// I programmatically determined that line 409 is the beginning of the loop in my input, I hardcoded this, although it could be determined by one additional iteration prior to this one
		if i == 409 {
			saveLoop = true
		}

		if saveLoop {
			loop = append(loop, i)
		}

		// If key i is present in the map, the instruction has already executed once, break and print the instruction. If key i is not present, add it to the map
		// Running this once and printing out the value is how I determined the hardcoded value above, which is still printed here
		_, ok := visitedLines[i]
		if ok {
			fmt.Printf("Instruction executed twice: %d | %s\n", i, dataArray[i])
			break
		}
		visitedLines[i] = true

		// We don't need to handle acc or nop instructions in this iteration

		// jmp instruction: add the number to i to move to the correct instruction (subtract one since the loop will also iterate it once)
		if strings.Contains(line, "jmp") {
			re := regexp.MustCompile(`[\+|\-]\d*`)
			numStr := re.FindString(line)
			num, _ := strconv.Atoi(numStr)
			i = i + num - 1
			continue
		}

	}

	// For each of the lines in the loop, iterate through the instructions again and change line loop[k] if it is a jmp or nop instruction to see if we reach the end of the document
	// Once the end of the document is reached, store which line was changed (loop[k])
	lineToChange := 0
	for k := range loop {
		if lineToChange != 0 {
			break
		}
		for k := range visitedLines {
			delete(visitedLines, k)
		}
		for i := 0; i <= len(dataArray); i++ {

			if i == len(dataArray) {
				lineToChange = loop[k]
				break
			}

			line := dataArray[i]
	
			// If key i is present in the map, the instruction has already executed once. If key i is not present, add it to the map
			_, ok := visitedLines[i]
			if ok {
				break
			}
			visitedLines[i] = true
	
			if i == loop[k] {
				if strings.Contains(line, "jmp") {
					re := regexp.MustCompile(`[\+|\-]\d*`)
					numStr := re.FindString(line)
					line = "nop " + numStr
				} else if strings.Contains(line, "nop") {
					re := regexp.MustCompile(`[\+|\-]\d*`)
					numStr := re.FindString(line)
					line = "jmp " + numStr
				}
			}

			// We don't need to handle acc or nop instructions in this iteration
	
			// jmp instruction: add the number to i to move to the correct instruction (subtract one since the loop will also iterate it once)
			if strings.Contains(line, "jmp") {
				re := regexp.MustCompile(`[\+|\-]\d*`)
				numStr := re.FindString(line)
				num, _ := strconv.Atoi(numStr)
				i = i + num - 1
				continue
			}
		}
	}

	fmt.Printf("Line to change: %d\n", lineToChange)
	acc := 0

	// One final iteration, making the proper change to fix the instructions and reach the end of the document, calculating and printing the acc output this time
	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if i == len(dataArray)-1 {
			fmt.Printf("End of document reached, acc = %d\n", acc)
			break
		}

		if i == lineToChange {
			if strings.Contains(line, "jmp") {
				re := regexp.MustCompile(`[\+|\-]\d*`)
				numStr := re.FindString(line)
				line = "nop " + numStr
			} else if strings.Contains(line, "nop") {
				re := regexp.MustCompile(`[\+|\-]\d*`)
				numStr := re.FindString(line)
				line = "jmp " + numStr
			}
		}

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
}