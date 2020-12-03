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

	var (
		x, y  int
		trees int
	)

	println("board is w", len(board[0]), "x h", len(board))
	for y <= len(board)-1 {
		if x >= len(board[0]) {
			x -= len(board[0])
		}

		println("checking x", x, "y", y)
		loc := string(board[y][x])
		if loc == "#" {
			trees++
		}

		x += 3 // move right 3...
		y++    // ..and down 1
	}

	println("encountered", trees, "trees")
}
