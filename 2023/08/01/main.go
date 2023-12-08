package main

import (
	"bufio"
	"os"
	"strings"

	ustrings "github.com/kaperys/advent-of-code/aoc/strings"
)

func main() {
	input, err := os.Open("2023/08/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var (
		directions string
		nodes      map[string][]string = make(map[string][]string)
	)

	scanner := bufio.NewScanner(input)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		if i == 0 {
			directions = strings.NewReplacer("L", "0", "R", "1").Replace(text)
		} else {
			if text != "" {
				nodes[text[0:3]] = []string{text[7:10], text[12:15]}
			}
		}
	}

	var (
		location string = "AAA"
		steps    int
	)

	for location != "ZZZ" {
		dir := ustrings.ToInt(string(directions[steps%len(directions)]))
		location = nodes[location][dir]
		steps++
	}

	println(steps)
}
