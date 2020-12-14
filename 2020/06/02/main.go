package main

import (
	"bufio"
	"os"
)

// https://adventofcode.com/2020/day/6#part2
func main() {
	input, err := os.Open("2020/06/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	type Group struct {
		People  int
		Answers map[string]int
	}

	var (
		groups       []Group
		currentGroup Group
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// empty line means new group
			groups = append(groups, currentGroup)
			currentGroup = Group{}
			continue
		}

		if currentGroup.Answers == nil {
			currentGroup.Answers = make(map[string]int)
		}

		currentGroup.People++
		for _, answer := range line {
			currentGroup.Answers[string(answer)]++
		}
	}

	groups = append(groups, currentGroup)

	var sum int
	for _, group := range groups {
		var valid int

		for _, answer := range group.Answers {
			if answer == group.People {
				valid++
			}
		}

		sum += valid
	}

	println("the sum of answers is", sum)
}
