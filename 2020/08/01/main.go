package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("08/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(strings.TrimSpace(string(input)), "\n")

	var i, a int
	visited := make(map[int]struct{})

	for {
		if _, ok := visited[i]; ok {
			println("accumulator is", a)
			return
		}

		visited[i] = struct{}{}

		var (
			operation string
			value     int
		)

		_, err := fmt.Sscanf(instructions[i], "%s %d", &operation, &value)
		if err != nil {
			panic(err)
		}

		switch operation {
		case "jmp":
			i += value
		case "acc":
			a += value
			i++
		case "nop":
			i++
		}
	}
}
