package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.Open("2023/04/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var total int

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		parts := strings.Split(text, ":")
		nums := strings.Split(parts[1], "|")

		winning := strings.Fields(nums[0])
		numbers := strings.Fields(nums[1])

		var score int
		for _, winner := range winning {
			if slices.Contains(numbers, winner) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score
	}

	println(total)
}
