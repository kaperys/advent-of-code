package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/3#part2
func main() {
	input, err := os.Open("2021/03/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var numbers []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	oxygenNumbers := make([]string, len(numbers))
	copy(oxygenNumbers, numbers)

	co2Numbers := make([]string, len(numbers))
	copy(co2Numbers, numbers)

	oxygen := oxygenGeneratorRating(oxygenNumbers, 0, "1")
	co2 := co2ScrubberRating(co2Numbers, 0, "0")

	oxygenDecimal, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		panic(err)
	}

	co2Decimal, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}

	println(oxygenDecimal * co2Decimal)
}

func oxygenGeneratorRating(numbers []string, position int, tie string) string {
	needle, _ := mostCommon(numbers, 0, "1")
	for {
		numbers = filterNumbers(numbers, needle, position)
		if len(numbers) == 1 {
			return numbers[0]
		}

		needle, _ = mostCommon(numbers, position+1, tie)
		position++
	}
}

func co2ScrubberRating(numbers []string, position int, tie string) string {
	needle, _ := mostCommon(numbers, 0, "0")

	// invert the needle
	if needle == "1" {
		needle = "0"
	} else {
		needle = "1"
	}

	for {
		numbers = filterNumbers(numbers, needle, position)
		if len(numbers) == 1 {
			return numbers[0]
		}

		var tied bool
		needle, tied = mostCommon(numbers, position+1, tie)

		if !tied {
			// invert the needle
			if needle == "1" {
				needle = "0"
			} else {
				needle = "1"
			}
		}

		position++
	}
}

func mostCommon(numbers []string, position int, tie string) (string, bool) {
	var i1, i0 int
	for _, number := range numbers {
		if string(number[position]) == "1" {
			i1++
		} else {
			i0++
		}
	}

	switch {
	case i1 == i0:
		return tie, true
	case i1 > i0:
		return "1", false
	case i0 > i1:
		return "0", false
	}

	return "-1", false
}

func filterNumbers(numbers []string, needle string, position int) []string {
	var filteredNumbers []string

	for _, number := range numbers {
		if string(number[position]) == needle {
			filteredNumbers = append(filteredNumbers, number)
		}
	}

	return filteredNumbers
}
