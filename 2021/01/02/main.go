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

	var depths []int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		depth, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		depths = append(depths, depth)
	}

	var last, larger int64
	for i := 0; i < len(depths); i++ {
		if i+2 > len(depths)-1 {
			break
		}

		x, y, z := depths[i], depths[i+1], depths[i+2]
		depth := x + y + z
		if depth > last {
			larger++
		}

		last = depth
	}

	println(larger - 1)
}
