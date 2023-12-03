package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type xy struct{ x, y int }

func main() {
	input, err := os.Open("2023/03/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		total int64
		grid  [][]rune
	)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		grid = append(grid, []rune(scanner.Text()))
	}

	for x, row := range grid {
		var (
			currentNumber     string
			hasAdjacentSymbol bool
		)

		for y, v := range row {
			isDigit := unicode.IsDigit(v)
			if isDigit {
				currentNumber += string(v)

				for _, neighbor := range neighborCoordinates(x, y) {
					if (neighbor.x >= 0 && neighbor.x <= len(row)-1) &&
						(neighbor.y >= 0 && neighbor.y <= len(grid)-1) {
						if isSymbol(grid[neighbor.x][neighbor.y]) {
							hasAdjacentSymbol = true
						}
					}
				}
			}

			if currentNumber != "" && // we've found some digits, and
				(!isDigit || y == len(row)-1) { // the current char isn't a digit, or we're at the end of a line
				if hasAdjacentSymbol { // and we've found an adjacent symbol
					partNumber, _ := strconv.ParseInt(currentNumber, 10, 64)
					total += partNumber
				}

				currentNumber = ""
				hasAdjacentSymbol = false
			}
		}
	}

	println(total)
}

func isSymbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func neighborCoordinates(x, y int) []xy {
	return []xy{
		{x: x, y: y - 1},
		{x: x + 1, y: y - 1},
		{x: x + 1, y: y},
		{x: x + 1, y: y + 1},
		{x: x, y: y + 1},
		{x: x - 1, y: y + 1},
		{x: x - 1, y: y},
		{x: x - 1, y: y - 1},
	}
}
