package main

import (
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/3#part2
func main() {
	input, err := ioutil.ReadFile("2020/03/input.txt")
	if err != nil {
		panic(err)
	}

	board := strings.Split(strings.TrimSpace(string(input)), "\n")

	traverse := func(right, down int) int {
		var (
			x, y  int
			trees int
		)

		for y <= len(board)-1 {
			if x >= len(board[0]) {
				x -= len(board[0])
			}

			if string(board[y][x]) == "#" {
				trees++
			}

			x += right
			y += down
		}

		return trees
	}

	totalTrees := 1
	for _, pos := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		totalTrees *= traverse(pos[0], pos[1])
	}

	println("encountered", totalTrees, "trees")
}
