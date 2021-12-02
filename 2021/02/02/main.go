package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2021/day/2#part2
func main() {
	input, err := os.Open("2021/02/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var depth, position, aim int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			command string
			unit    int64
		)

		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &command, &unit)
		if err != nil {
			panic(err)
		}

		switch strings.TrimSpace(strings.ToLower(command)) {
		case "forward":
			position += unit
			depth += (aim * unit)
		case "down":
			aim += unit
		case "up":
			aim -= unit
		}
	}

	println(depth * position)
}
