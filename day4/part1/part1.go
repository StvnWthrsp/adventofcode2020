/*
Steven Weatherspoon
2020 Advent of Code Challenge
Day 4, Part 1

This program takes the first argument as the input filepath
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
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
	// Since the logic will check for the end of a passport using a blank line, we must append an empty string to the end of the array
	dataArray = append(dataArray, "")

	validPassportCount := 0

	// Create a hash map to store a boolean based whether the given field exists
	fields := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false}

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if len(line) != 0 {
			// If the line is not an empty string, check for existence of the needed fields
			// Since a single passport can span multiple lines, we have to make sure they are each false, so that we don't set a true value from a previous line back to false
			if !fields["byr"] {fields["byr"] = strings.Contains(line, "byr")}
			if !fields["iyr"] {fields["iyr"] = strings.Contains(line, "iyr")}
			if !fields["eyr"] {fields["eyr"] = strings.Contains(line, "eyr")}
			if !fields["hgt"] {fields["hgt"] = strings.Contains(line, "hgt")}
			if !fields["hcl"] {fields["hcl"] = strings.Contains(line, "hcl")}
			if !fields["ecl"] {fields["ecl"] = strings.Contains(line, "ecl")}
			if !fields["pid"] {fields["pid"] = strings.Contains(line, "pid")}
			//if !fields["cid"] {fields["cid"] = strings.Contains(line, "cid")}
		} else {
			// If the line is an empty string, iterate through the map to check for any false values. If any value is false, the passport is invalid.
			valid := true
			for k := range fields {
				if !fields[k] {
					valid = false
				}
			}
			// Set each value back to false
			for k := range fields {
				fields[k] = false
			}
			if valid {
				validPassportCount++
			}
		}
	}
	fmt.Printf("Valid passport count: %d\n", validPassportCount)
}