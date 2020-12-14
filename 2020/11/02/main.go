package main

import (
	"image"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/11#part2
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

	seats := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: len(in[0]), Y: len(in)},
	}

	for y, row := range in {
		for x, seat := range row {
			var occupied int

			adjacent := []image.Point{
				{X: -1, Y: -1}, // north west
				{Y: -1},        // north
				{X: +1, Y: -1}, // north east
				{X: -1},        // west
				{X: +1},        // east
				{X: -1, Y: +1}, // south west
				{Y: +1},        // south
				{X: +1, Y: +1}, // south east
			}

			start := image.Point{X: x, Y: y}
			for _, dir := range adjacent {
				loc := start.Add(dir)

				for loc.In(seats) {
					if in[loc.Y][loc.X] == 'L' {
						break
					}

					if in[loc.Y][loc.X] == '#' {
						occupied++
						break
					}

					loc = loc.Add(dir)
				}
			}

			if seat == 'L' && occupied == 0 {
				out[y] += "#"
			} else if seat == '#' && occupied >= 5 {
				out[y] += "L"
			} else {
				out[y] += string(seat)
			}
		}
	}

	return out
}
