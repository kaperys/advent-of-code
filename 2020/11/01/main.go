package main

import (
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/11
func main() {
	input, err := ioutil.ReadFile("2020/11/input.txt")
	if err != nil {
		panic(err)
	}

	in := strings.Split(strings.TrimSpace(string(input)), "\n")

	for {
		out := run(in)

		o := strings.Join(out, "")
		i := strings.Join(in, "")

		if o != i {
			in = out
			continue
		}

		println("there are", strings.Count(i, "#"), "occupied seats")
		return
	}
}

func run(in []string) []string {
	out := make([]string, len(in))

	for y, row := range in {
		for x, seat := range row {
			var occupied int

			adjacent := [][]int{
				{x - 1, y - 1}, // north west
				{x, y - 1},     // north
				{x + 1, y - 1}, // north east
				{x - 1, y},     // west
				{x + 1, y},     // east
				{x - 1, y + 1}, // south west
				{x, y + 1},     // south
				{x + 1, y + 1}, // south east
			}

			for _, xy := range adjacent {
				ax, ay := xy[0], xy[1]

				if ax < 0 || ay < 0 {
					continue
				}

				if ax >= len(row) || ay >= len(in) {
					continue
				}

				if in[ay][ax] == '#' {
					occupied++
				}
			}

			if seat == 'L' && occupied == 0 {
				out[y] += "#"
			} else if seat == '#' && occupied >= 4 {
				out[y] += "L"
			} else {
				out[y] += string(seat)
			}
		}
	}

	return out
}
