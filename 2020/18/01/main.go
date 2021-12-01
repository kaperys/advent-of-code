package main

import (
	"bufio"
	"os"
	"strings"
)

// https://adventofcode.com/2020/day/18
func main() {
	input, err := os.Open("2020/18/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var operands, operators string

		for _, l := range scanner.Text() {
			switch strings.TrimSpace(string(l)) {
			case "":
				continue
			case "+", "*":
				operators += string(l)
			default:
				operands += string(l)
			}
		}

		rpn := operands + operators
		println(rpn)
	}
}
