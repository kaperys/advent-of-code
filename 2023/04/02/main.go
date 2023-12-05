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

	var (
		total int

		counts map[int]int = make(map[int]int)
		cards  []string
	)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		counts[i] = 1 // start with 1 of each card
		cards = append(cards, scanner.Text())
	}

	for i, card := range cards {
		nums := strings.Split(strings.Split(card, ":")[1], "|")

		winning := strings.Fields(nums[0])
		numbers := strings.Fields(nums[1])

		var score int
		for _, winner := range winning {
			if slices.Contains(numbers, winner) {
				score++
			}
		}

		if score > 0 {
			for j := i + 1; j <= i+score; j++ {
				counts[j] += counts[i]
			}
		}
	}

	for _, v := range counts {
		total += v
	}

	println(total)
}
