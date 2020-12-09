package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open("2020/05/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var max int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rows := line[0:7]
		columns := line[7:10]

		row := search(0, 127, "F", "B", rows)
		column := search(0, 7, "L", "R", columns)
		seatID := (row * 8) + column

		if seatID > max {
			max = seatID
		}
	}

	println("the highest seat ID is", max)
}

func search(min, max int, low, high, input string) int {
	for _, dir := range input {
		mid := (min + max) / 2

		move := string(dir)
		if move == low {
			max = mid
		}

		if move == high {
			min = mid + 1
		}

		if min == max {
			return max
		}
	}

	return -1
}
