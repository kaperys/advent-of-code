package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1
func main() {
	input, err := os.Open("2021/01/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var last, larger int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		depth, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		if depth > last {
			larger++
		}

		last = depth
	}

	println(larger - 1)
}
