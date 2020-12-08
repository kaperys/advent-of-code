package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("04/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		passports       []map[string]string
		currentPassport = make(map[string]string)
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line means new passport
			passports = append(passports, currentPassport)
			currentPassport = make(map[string]string)
			continue
		}

		components := strings.Split(line, " ")
		for _, component := range components {
			parts := strings.Split(component, ":")
			currentPassport[parts[0]] = parts[1]
		}
	}

	passports = append(passports, currentPassport)

	var validPassports int

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" /*"cid"*/}
	for _, passport := range passports {
		valid := true

		for _, field := range requiredFields {
			if _, ok := passport[field]; !ok {
				valid = false
			}
		}

		if valid {
			validPassports++
		}
	}

	println("there are", validPassports, "valid passports")
}
