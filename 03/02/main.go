package main

import (
	"bufio"
	"os"
)

func main() {
	input, err := os.Open("03/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var board []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	println("board is w", len(board[0]), "x h", len(board))
	traverse := func(right, down int) int {
		var (
			x, y  int
			trees int
		)

		for y <= len(board)-1 {
			if x >= len(board[0]) {
				x -= len(board[0])
			}

			// println("checking x", x, "y", y)
			loc := string(board[y][x])
			if loc == "#" {
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
