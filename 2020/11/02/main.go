package main

import (
	"image"
	"io/ioutil"
	"strings"
)

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

		// TODO(kaperys) implement a less brute-forcey way to do this
		var occupied int
		for _, s := range i {
			if s == '#' {
				occupied++
			}
		}

		println("there are", occupied, "occupied seats")
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

			adjacent := []func(image.Point) image.Point{
				func(point image.Point) image.Point { // north west
					return point.Add(image.Point{X: -1, Y: -1})
				},
				func(point image.Point) image.Point { // north
					return point.Add(image.Point{Y: -1})
				},
				func(point image.Point) image.Point { // north east
					return point.Add(image.Point{X: +1, Y: -1})
				},
				func(point image.Point) image.Point { // west
					return point.Add(image.Point{X: -1})
				},
				func(point image.Point) image.Point { // east
					return point.Add(image.Point{X: +1})
				},
				func(point image.Point) image.Point { // south west
					return point.Add(image.Point{X: -1, Y: +1})
				},
				func(point image.Point) image.Point { // south
					return point.Add(image.Point{Y: +1})
				},
				func(point image.Point) image.Point { // south east
					return point.Add(image.Point{X: +1, Y: +1})
				},
			}

			start := image.Point{X: x, Y: y}
			for _, fn := range adjacent {
				loc := fn(start)

				for loc.In(seats) {
					if in[loc.Y][loc.X] == 'L' {
						break
					}

					if in[loc.Y][loc.X] == '#' {
						occupied++
						break
					}

					loc = fn(loc)
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
