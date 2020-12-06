package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open("06/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		groups       []map[string]bool
		currentGroup = make(map[string]bool)
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line means new group
			groups = append(groups, currentGroup)
			currentGroup = make(map[string]bool)
			continue
		}

		for _, answer := range line {
			currentGroup[string(answer)] = true
		}
	}

	groups = append(groups, currentGroup)

	var sum int
	for _, group := range groups {
		sum += len(group)
	}

	println("the sum of answers is", sum)
}
