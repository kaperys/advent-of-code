package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var rules = map[string]int64{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	input, err := os.Open("2023/02/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var total int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ":")

		game, _ := strconv.ParseInt(strings.TrimSpace(parts[0])[5:], 10, 64)
		turns := strings.Split(parts[1], ";")

		if isPossible(turns) {
			total += game
		}
	}

	println(total)
}

func isPossible(turns []string) bool {
	for _, turn := range turns {
		cubes := strings.Split(turn, ",")
		for _, cube := range cubes {
			parts := strings.Split(strings.TrimSpace(cube), " ")

			count, _ := strconv.ParseInt(parts[0], 10, 64)
			colour := parts[1]

			rule, ok := rules[colour]
			if ok && count > rule {
				return false
			}
		}
	}

	return true
}
