package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/03/input.txt")
	if err != nil {
		panic(err)
	}

	board := strings.Split(strings.TrimSpace(string(input)), "\n")

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

		x += 3 // move right 3...
		y++    // ..and down 1
	}

	println("encountered", trees, "trees")
}
