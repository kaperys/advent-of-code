package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/13
func main() {
	input, err := ioutil.ReadFile("2020/13/input.txt")
	if err != nil {
		panic(err)
	}

	in := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		earliestBus int
		earliestDel int = math.MaxInt64
	)

	timestamp, _ := strconv.Atoi(in[0])
	for _, bus := range strings.Split(in[1], ",") {
		if bus == "x" {
			continue
		}

		ts, _ := strconv.Atoi(bus)
		earliest := ((timestamp / ts) * ts) + ts

		if delta := (timestamp - earliest) * -1; delta < earliestDel {
			earliestDel = delta
			earliestBus = ts
		}
	}

	println("I need to wait", earliestBus*earliestDel, "to take bus", earliestBus)
}
