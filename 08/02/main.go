package main

import (
	"bufio"
	"os"
	"strconv"
)

type Instruction struct {
	Operation string
	Value     int
}

func main() {
	input, err := os.Open("08/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var instructions []Instruction

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line[4:])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{
			Operation: line[0:3],
			Value:     value,
		})
	}

	for i := 0; i < len(instructions); i++ {
		modifiedInstructions := make([]Instruction, len(instructions))
		copy(modifiedInstructions, instructions)

		switch modifiedInstructions[i].Operation {
		case "jmp":
			modifiedInstructions[i].Operation = "nop"
		case "nop":
			modifiedInstructions[i].Operation = "jmp"
		}

		j, a := run(modifiedInstructions)
		if j >= len(modifiedInstructions) {
			println("accumulator is", a)
			return
		}
	}
}

func run(instructions []Instruction) (int, int) {
	var (
		i           int
		accumulator int

		visited = make(map[int]struct{})
	)

	for {
		if i >= len(instructions) {
			return i, accumulator
		}

		if _, ok := visited[i]; ok {
			return i, accumulator
		}

		op := instructions[i]
		switch op.Operation {
		case "jmp":
			visited[i] = struct{}{}
			i += op.Value

		case "acc":
			visited[i] = struct{}{}
			accumulator += op.Value

			i++
		case "nop":
			visited[i] = struct{}{}
			i++
		}
	}
}
