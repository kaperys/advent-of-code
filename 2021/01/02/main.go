package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1#part2
func main() {
	input, err := os.Open("2021/01/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var prev1, prev2, lastDepth, count int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		current, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		if prev1 != 0 && prev2 != 0 {
			depth := current + prev1 + prev2
			if depth > lastDepth {
				count++
			}

			lastDepth = depth
		}

		prev2 = prev1
		prev1 = current
	}

	println(count - 1)
}
