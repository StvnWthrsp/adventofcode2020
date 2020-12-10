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
	"regexp"
	"strconv"
)

func validatePassport(passport map[string]string) bool {
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}

	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}

	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}

	hgt := passport["hgt"]
	hcl := passport["hcl"]
	ecl := passport["ecl"]
	pid := passport["pid"]

	if byr < 1920 || byr > 2002 {
		return false
	}

	if iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr < 2020 || eyr > 2030 {
		return false
	}

	if len(hgt) == 0 {
		return false
	}
	hgtUnit := hgt[len(hgt)-2:]
	hgtDigit, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	switch (hgtUnit) {
	case "cm":
		if hgtDigit < 150 || hgtDigit > 193 {
			return false
		}
	case "in":
		if hgtDigit < 59 || hgtDigit > 76 {
			return false
		}
	default:
		return false
	}

	re := regexp.MustCompile(`^#[a-z0-9_]{6}$`)
	if !re.MatchString(hcl) {
		return false
	}

	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	re = regexp.MustCompile(`^\d{9}$`)
	if !re.MatchString(pid) {
		return false
	}

	return true
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
	// Since the logic will check for the end of a passport using a blank line, we must append an empty string to the end of the array
	dataArray = append(dataArray, "")

	currentPassport := make(map[string]string)
	re := regexp.MustCompile(`(\S{3}:\S+)`)
	validPassports := 0

	for i := 0; i < len(dataArray); i++ {
		line := dataArray[i]

		if len(line) == 0 {
			// At each blank line, pass current passport information to the validation function, then clear the map
			isValid := validatePassport(currentPassport)
			if isValid {
				validPassports++
			}
			for key := range currentPassport {
				delete(currentPassport, key)
			}
			continue
		}
		// If the line is not blank, use regex to get each key/value chunk, then split it into a variable for each
		for _, chunk := range re.FindAllString(line, -1) {
			splitted := strings.Split(chunk, ":")
			code, val := splitted[0], splitted[1]
			currentPassport[code] = val;
		}
	}
	fmt.Printf("Valid passports: %d\n", validPassports)
}