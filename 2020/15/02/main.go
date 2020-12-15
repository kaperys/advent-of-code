package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const target = 30000000

// https://adventofcode.com/2020/day/15#part2
func main() {
	input, err := ioutil.ReadFile("2020/15/input.txt")
	if err != nil {
		panic(err)
	}

	var numbers []int

	for _, i := range strings.Split(strings.TrimSpace(string(input)), ",") {
		i, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, i)
	}

	var (
		last   int
		memory = make(map[int]int)
	)

	for i := 0; i < target; i++ {
		if i < len(numbers) {
			memory[numbers[i]] = i + 1
			last = numbers[i]
			continue
		}

		l, ok := memory[last]
		if ok {
			memory[last] = i
			last = i - l
			continue
		}

		memory[last] = i
		last = 0
	}

	println("the last number spoken was", last)
}
