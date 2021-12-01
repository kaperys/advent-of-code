package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/16#part2
func main() {
	input, err := os.Open("2020/16/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		patterns  [][]int
		errorRate int

		readPatterns, readTickets bool
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if !readPatterns {
			if line == "" {
				readPatterns = true
				continue
			}

			parts := strings.Split(line, ":")
			for _, pattern := range strings.Split(strings.TrimSpace(parts[1]), " or ") {

				parts := strings.Split(pattern, "-")
				min, _ := strconv.Atoi(parts[0])
				max, _ := strconv.Atoi(parts[1])

				patterns = append(patterns, []int{min, max})
			}
		}

		if !readTickets {
			if line == "nearby tickets:" {
				readTickets = true
				continue
			}
		}

		if readTickets {
			for _, value := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(value)

				var valid bool
				for _, pattern := range patterns {
					if value >= pattern[0] && value <= pattern[1] {
						valid = true
						break
					}
				}

				if !valid {
					errorRate += value
				}
			}
		}
	}

	println("the scanning error rate is", errorRate)
}
