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
		grid  []string
	)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		grid = append(grid, scanner.Text())
	}

	visited := make(map[xy]bool)
	for x, row := range grid {
		for y, v := range row {
			if v == '*' {
				var gears []int64

				for _, neighbor := range neighborCoordinates(x, y) {
					if (neighbor.x >= 0 && neighbor.x <= len(row)-1) &&
						(neighbor.y >= 0 && neighbor.y <= len(grid)-1) {
						if unicode.IsDigit(rune(grid[neighbor.x][neighbor.y])) && !visited[neighbor] {
							start, end := neighbor.y, neighbor.y

							for start > 0 && unicode.IsDigit(rune(grid[neighbor.x][start-1])) {
								visited[xy{x: neighbor.x, y: start - 1}] = true
								start--
							}

							for end < len(grid[0])-1 && unicode.IsDigit(rune(grid[neighbor.x][end+1])) {
								visited[xy{x: neighbor.x, y: end + 1}] = true
								end++
							}

							gear, _ := strconv.ParseInt(string(grid[neighbor.x][start:end+1]), 10, 64)
							gears = append(gears, gear)
						}
					}
				}

				if len(gears) > 1 {
					var gearRatio int64 = 1

					for _, gear := range gears {
						gearRatio *= gear
					}

					total += int64(gearRatio)
				}
			}
		}
	}

	println(total)
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
