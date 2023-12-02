package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var cubesPattern = regexp.MustCompile(`([0-9]+) (red|blue|green)+`)

func main() {
	input, err := os.Open("2023/02/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var total int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cubes := cubesPattern.FindAllStringSubmatch(scanner.Text(), -1)
		count := map[string]int64{"green": 0, "blue": 0, "red": 0}

		for _, cube := range cubes {
			num, _ := strconv.ParseInt(cube[1], 10, 64)

			if num > count[cube[2]] {
				count[cube[2]] = num
			}
		}

		total += count["red"] * count["green"] * count["blue"]
	}

	println(total)
}
