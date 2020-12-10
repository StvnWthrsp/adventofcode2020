/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 7, Part 2

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

// Recursively search the list to calculate the total number of bags in the shiny gold bag of bags
func findBagsAbove(dataArray []string, bagToFind string, runningTotal int, multiplier int) int {
	currentTotal := runningTotal

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if strings.Contains(line, "no other bags") {
			continue
		}

		if strings.Contains(line, bagToFind) {
			// Use regexp to get the first bag name on the line
			re := regexp.MustCompile(`^\w+ \w+ \w{3}`)
			bagName := re.FindString(line)

			// Use regexp to get each of the inner bags
			re = regexp.MustCompile(`\d \w+ \w+ \w{3}`)
			innerBags := re.FindAllString(line, -1)

			// If the bagName on line i matches the given bagToFind, calculate the number of each kind of bags inside of it, as well as the number of bags inside those bags
			if bagName == bagToFind {
				for _, e := range innerBags {
					// Use regexp to separate the line into a digit and bag name
					re = regexp.MustCompile(`[^ 0-9]+ \D+ \D{3}`)
					innerBagName := re.FindString(e)
					re = regexp.MustCompile(`\d`)
					digit := re.FindString(e)
					digitInt, _ := strconv.Atoi(digit)

					// Total number of bag innerBagName
					currentTotal += multiplier*digitInt
					// New multiplier for any bags inside of innerBagName
					newMultiplier := multiplier*digitInt

					// Recursively call the function again to get number of bags in inner bag innerBagName
					currentTotal = findBagsAbove(dataArray, innerBagName, currentTotal, newMultiplier)
				}
				continue
			}
		}
	}
	return currentTotal
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

	// Begin recursion by looking in the shiny gold bag with an initial total of 0 and multiplier of 1
	total := findBagsAbove(dataArray, "shiny gold bag", 0, 1)

	fmt.Printf("Total: %d\n", total)
}