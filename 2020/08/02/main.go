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

	for i := 0; i < len(instructions); i++ {
		modifiedInstructions := make([]string, len(instructions))
		copy(modifiedInstructions, instructions)

		modifiedInstructions[i] = strings.NewReplacer(
			"jmp", "nop",
			"nop", "jmp",
		).Replace(modifiedInstructions[i])

		j, a := run(modifiedInstructions)
		if j >= len(modifiedInstructions) {
			println("accumulator is", a)
			return
		}
	}
}

func run(instructions []string) (int, int) {
	var i, a int
	visited := make(map[int]struct{})

	for {
		if i >= len(instructions) {
			return i, a
		}

		if _, ok := visited[i]; ok {
			return i, a
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
