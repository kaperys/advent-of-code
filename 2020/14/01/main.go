package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/14
func main() {
	input, err := ioutil.ReadFile("2020/14/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		and int64
		or  int64
	)

	memory := make(map[int64]int64)
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if strings.HasPrefix(line, "mask") {
			and, _ = strconv.ParseInt(strings.ReplaceAll(line[7:], "X", "1"), 2, 0)
			or, _ = strconv.ParseInt(strings.ReplaceAll(line[7:], "X", "0"), 2, 0)
			continue
		}

		var position, value int64
		fmt.Sscanf(line, "mem[%d] = %d", &position, &value)

		memory[position] = value&int64(and) | int64(or)
	}

	var sum int64
	for _, v := range memory {
		sum += v
	}

	println("the sum of memory values is", sum)
}
