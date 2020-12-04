package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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
			val, ok := passport[field]
			if !ok {
				// println(field, "not found")
				valid = false
				break
			}

			// byr (Birth Year) - four digits; at least 1920 and at most 2002
			if field == "byr" {
				if len(val) != 4 {
					println(field, "invalid len", val)
					valid = false
					break
				}

				byr, err := strconv.Atoi(val)
				if err != nil {
					println(field, "invalid strconv", val)
					valid = false
					break
				}

				if !(byr >= 1920 && byr <= 2002) {
					println(field, "invalid range", val)
					valid = false
					break
				}
			}

			// iyr (Issue Year) - four digits; at least 2010 and at most 2020
			if field == "iyr" {
				if len(val) != 4 {
					println(field, "invalid len", val)
					valid = false
					break
				}

				iyr, err := strconv.Atoi(val)
				if err != nil {
					println(field, "invalid strconv", val)
					valid = false
					break
				}

				if !(iyr >= 2010 && iyr <= 2020) {
					println(field, "invalid range", val)
					valid = false
					break
				}
			}

			// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
			if field == "eyr" {
				if len(val) != 4 {
					println(field, "invalid len", val)
					valid = false
					break
				}

				eyr, err := strconv.Atoi(val)
				if err != nil {
					println(field, "invalid strconv", val)
					valid = false
					break
				}

				if !(eyr >= 2020 && eyr <= 2030) {
					println(field, "invalid range", val)
					valid = false
					break
				}
			}

			// hgt (Height) - a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			if field == "hgt" {
				switch {
				case strings.HasSuffix(val, "cm"):
					hgt, err := strconv.Atoi(val[:len(val)-2])
					if err != nil {
						println(field, "invalid strconv", val)
						valid = false
						break
					}

					if !(hgt >= 150 && hgt <= 193) {
						println(field, "invalid range", val)
						valid = false
						break
					}
				case strings.HasSuffix(val, "in"):
					hgt, err := strconv.Atoi(val[:len(val)-2])
					if err != nil {
						println(field, "invalid strconv")
						valid = false
						break
					}

					if !(hgt >= 59 && hgt <= 76) {
						println(field, "invalid range", val)
						valid = false
						break
					}
				default:
					println(field, "invalid suffix", val)
					valid = false
					break
				}
			}

			// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
			if field == "hcl" {
				if len(val) != 7 {
					println(field, "invalid len", val)
					valid = false
					break
				}

				match, _ := regexp.MatchString("^#(?:[0-9a-f]{6})$", val)
				if !match {
					println(field, "invalid regex", val)
					valid = false
					break
				}
			}

			// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth
			if field == "ecl" {
				switch val {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				default:
					println(field, "invalid string", val)
					valid = false
					break
				}
			}

			// pid (Passport ID) - a nine-digit number, including leading zeroes
			if field == "pid" {
				if len(val) != 9 {
					println(field, "invalid len", val)
					valid = false
					break
				}

				_, err := strconv.Atoi(val)
				if err != nil {
					println(field, "invalid strconv")
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
