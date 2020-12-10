/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 7, Part 1

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
)

// Recursively search the list to find all bags that can ultimately contain the given "bagToFind"
func findBagsAbove(dataArray []string, goldBags map[string]bool, bagToFind string) map[string]bool {
	outputMap := goldBags
	newBagAdded := false
	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if strings.Contains(line, "no other bags") {
			continue
		}

		if strings.Contains(line, bagToFind) {
			// Use regexp to get the first bag name on the line
			re := regexp.MustCompile(`^\w+ \w+ \w{3}`)
			bagName := re.FindString(line)
			// Required to exclude the rule pertaining to the bagToFind itself
			if bagName == bagToFind {
				continue
			}
			// If the bagName is not in the map, add it and set newBagAdded to true to enable recursion
			_, ok := outputMap[bagName]
			if !ok {
				newBagAdded = true
				outputMap[bagName] = true
			}
			continue
		}
	}
	// If a new bag was added to the map, we have to check again to see if any other bag can hold the new bag. Rescursion will stop when no new bags are added to the map.
	if newBagAdded {
		for k := range outputMap {
			outputMap = findBagsAbove(dataArray, outputMap, k)
		}
	}
	return outputMap
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

	total := 0
	goldBags := make(map[string]bool)

	// Begin recursion with the bag to find, shiny gold bag
	goldBags = findBagsAbove(dataArray, goldBags, "shiny gold bag")

	// Add up the keys in the map to count the total number of bags
	for k := range goldBags {
		fmt.Printf("%s\n", k)
		total++
	}
	fmt.Printf("Total: %d\n", total)
}