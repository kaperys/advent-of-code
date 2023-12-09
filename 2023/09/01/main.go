package main

import (
	"bufio"
	"os"

	ustrings "github.com/kaperys/advent-of-code/aoc/strings"
)

func main() {
	input, err := os.Open("2023/09/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var total int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		in := ustrings.ToIntSlice(scanner.Text())
		var rows [][]int = [][]int{in}

		for r := 0; r < len(rows); r++ {
			var (
				deltas []int
				zeros  int
			)

			for i := 0; i < len(rows[r])-1; i++ {
				delta := rows[r][i+1] - rows[r][i]
				if delta == 0 {
					zeros++
				}

				deltas = append(deltas, delta)
			}

			rows = append(rows, deltas)
			if zeros == len(rows[r])-1 {
				break
			}
		}

		for r := len(rows) - 2; r > 0; r-- {
			new := rows[r][len(rows[r])-1] + rows[r-1][len(rows[r-1])-1]
			rows[r-1] = append(rows[r-1], new)
		}

		total += rows[0][len(rows[0])-1]
	}

	println(total)
}
