package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
	gamePattern  = regexp.MustCompile(`Game ([0-9]+):`)
	cubesPattern = regexp.MustCompile(`([0-9]+) (red|blue|green)+`)

	rules = map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

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

		game, _ := strconv.ParseInt(gamePattern.FindStringSubmatch(text)[1], 10, 64)
		cubes := cubesPattern.FindAllStringSubmatch(text, -1)

		isPossible := true
		for _, cube := range cubes {
			num, _ := strconv.ParseInt(cube[1], 10, 64)
			if num > rules[cube[2]] {
				isPossible = false
				break
			}
		}

		if isPossible {
			total += game
		}
	}

	println(total)
}
