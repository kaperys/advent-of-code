package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open("05/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	seats := make(map[int]bool)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rows := line[0:7]
		columns := line[7:10]

		row := search(0, 127, "F", "B", rows)
		column := search(0, 7, "L", "R", columns)
		seatID := (row * 8) + column

		// println(row)
		// println(column)
		// println(seatID)

		seats[seatID] = true
	}

	for i := 0; i < len(seats); i++ {
		if _, ok := seats[i]; !ok {
			_, before := seats[i-1]
			_, after := seats[i+1]

			if before && after {
				println("my seat is", i)
			}
		}
	}
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

		// println(min, max, ":", mid, move)
	}

	return -1
}
