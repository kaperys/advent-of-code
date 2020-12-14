package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/12/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		ship image.Point
		dir  int
	)

	for _, instruction := range instructions {
		ins := instruction[:1]
		val, _ := strconv.Atoi(instruction[1:])

		if ins == "F" {
			switch dir {
			case 0:
				ins = "E"
			case 1:
				ins = "S"
			case 2:
				ins = "W"
			case 3:
				ins = "N"
			}
		}

		switch ins {
		case "N":
			ship = ship.Add(image.Point{Y: val})
		case "S":
			ship = ship.Add(image.Point{Y: -val})
		case "E":
			ship = ship.Add(image.Point{X: val})
		case "W":
			ship = ship.Add(image.Point{X: -val})
		case "L":
			dir = 4 + (dir-(val/90))%4
			if dir >= 4 {
				dir -= 4
			}
		case "R":
			dir = (dir + (val / 90)) % 4
		}
	}

	fmt.Printf("ship is at x%d y%d: %.0f\n", ship.X, ship.Y, math.Abs(float64(ship.X))+math.Abs(float64(ship.Y)))
}
