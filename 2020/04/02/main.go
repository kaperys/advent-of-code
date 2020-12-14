package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/4#part2
func main() {
	input, err := os.Open("2020/04/input.txt")
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
			val, ok := passport[field]
			if !ok {
				valid = false
				break
			}

			switch field {
			case "byr": // byr (Birth Year) - four digits; at least 1920 and at most 2002
				match, _ := regexp.MatchString("^(19[2-9][0-9]|200[0-2])$", val)
				if !match {
					valid = false
					break
				}
			case "iyr": // iyr (Issue Year) - four digits; at least 2010 and at most 2020
				match, _ := regexp.MatchString("^(201[0-9]|2020)$", val)
				if !match {
					valid = false
					break
				}
			case "eyr": // eyr (Expiration Year) - four digits; at least 2020 and at most 2030
				match, _ := regexp.MatchString("^(20[2][0-9]|2030)$", val)
				if !match {
					valid = false
					break
				}
			case "hgt": // hgt (Height) - a number followed by either cm or in:
				// If cm, the number must be at least 150 and at most 193.
				// If in, the number must be at least 59 and at most 76.
				switch {
				case strings.HasSuffix(val, "cm"):
					hgt, _ := strconv.Atoi(val[:len(val)-2])
					if !(hgt >= 150 && hgt <= 193) {
						valid = false
						break
					}
				case strings.HasSuffix(val, "in"):
					hgt, _ := strconv.Atoi(val[:len(val)-2])
					if !(hgt >= 59 && hgt <= 76) {
						valid = false
						break
					}
				default:
					valid = false
					break
				}
			case "hcl": // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
				match, _ := regexp.MatchString("^#(?:[0-9a-f]{6})$", val)
				if !match {
					valid = false
					break
				}
			case "ecl": // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
				match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", val)
				if !match {
					valid = false
					break
				}
			case "pid": // pid (Passport ID) - a nine-digit number, including leading zeroes
				match, _ := regexp.MatchString("^([0-9]{9})$", val)
				if !match {
					valid = false
					break
				}
			}
		}

		if valid {
			validPassports++
		}
	}

	println("there are", validPassports, "valid passports")
}
