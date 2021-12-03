package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/3
func main() {
	input, err := os.Open("2021/03/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		bits = make([]int64, 12)
		len  int64
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		for i, pos := range scanner.Text() {
			if string(pos) == "1" {
				bits[i] += 1
			}
		}

		len++
	}

	var gamma, epsilon string
	for _, bit := range bits {
		if bit > len/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaDecimal, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonDecimal, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	println(gammaDecimal * epsilonDecimal)
}
