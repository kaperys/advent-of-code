package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

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

				// above and below
				if x != 0 { // above
					if isSymbol(grid[x-1][y]) {
						println("found symbol", string(grid[x-1][y]), "above", string(v))
						hasAdjacentSymbol = true
					}
				}

				if x < len(grid)-1 { // below
					if isSymbol(grid[x+1][y]) {
						println("found symbol", string(grid[x+1][y]), "below", string(v))
						hasAdjacentSymbol = true
					}
				}

				// left and right
				if y != 0 { // left
					if isSymbol(grid[x][y-1]) {
						println("found symbol", string(grid[x][y-1]), "left of", string(v))
						hasAdjacentSymbol = true
					}
				}

				if y != len(row)-1 { // right
					if isSymbol(grid[x][y+1]) {
						println("found symbol", string(grid[x][y+1]), "right of", string(v))
						hasAdjacentSymbol = true
					}
				}

				// diagonals
				if x != 0 { // above
					if y != 0 { // left
						if isSymbol(grid[x-1][y-1]) {
							println("found symbol", string(grid[x-1][y-1]), "above left of", string(v))
							hasAdjacentSymbol = true
						}
					}

					if y != len(row)-1 { // right
						if isSymbol(grid[x-1][y+1]) {
							println("found symbol", string(grid[x-1][y+1]), "above right of", string(v))
							hasAdjacentSymbol = true
						}
					}
				}

				if x < len(grid)-2 { // below
					if y != 0 { // left
						if isSymbol(grid[x+1][y-1]) {
							println("found symbol", string(grid[x+1][y-1]), "below left of", string(v))
							hasAdjacentSymbol = true
						}
					}

					if y != len(row)-1 { // right
						if isSymbol(grid[x+1][y+1]) {
							println("found symbol", string(grid[x+1][y+1]), "below right of", string(v))
							hasAdjacentSymbol = true
						}
					}
				}
			}

			if currentNumber != "" && // we've found some digits, and
				(!isDigit || y == len(row)-1) && // the current char isn't a digit, or we're at the end of a line
				hasAdjacentSymbol { // and we've found an adjacent symbol
				part, _ := strconv.ParseInt(currentNumber, 10, 64)
				total += part

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
