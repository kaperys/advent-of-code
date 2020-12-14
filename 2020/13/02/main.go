package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/13#part2
func main() {
	input, err := ioutil.ReadFile("2020/13/input.txt")
	if err != nil {
		panic(err)
	}

	in := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		timestamp int
		j         int = 1
	)

	for i, bus := range strings.Split(in[1], ",") {
		if bus == "x" {
			continue
		}

		ts, _ := strconv.Atoi(bus)
		for (timestamp+i)%ts != 0 {
			timestamp += j
		}

		j *= ts
	}

	println("the earliest timestamp is", timestamp)
}
