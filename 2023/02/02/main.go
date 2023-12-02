package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
		parts := strings.Split(text, ":")

		turns := strings.Split(parts[1], ";")
		power := minRequired(turns)

		total += power
	}

	println(total)
}

func minRequired(turns []string) int64 {
	colours := map[string]int64{"red": 0, "blue": 0, "green": 0}

	for _, turn := range turns {
		cubes := strings.Split(turn, ",")
		for _, cube := range cubes {
			parts := strings.Split(strings.TrimSpace(cube), " ")

			count, _ := strconv.ParseInt(parts[0], 10, 64)
			colour := parts[1]

			if colours[colour] < count {
				colours[colour] = count
			}
		}
	}

	return colours["green"] * colours["red"] * colours["blue"]
}
