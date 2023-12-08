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

	var paths []int
	for node := range nodes {
		if string(node[2]) == "A" {
			cur := node

			var steps int
			for string(cur[2]) != "Z" {
				dir := ustrings.ToInt(string(directions[steps%len(directions)]))
				cur = nodes[cur][dir]
				steps++
			}

			paths = append(paths, steps)
		}
	}

	println(lcm(paths...))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(x ...int) int {
	if len(x) == 1 {
		return x[0]
	} else if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	}

	return x[0] * x[1] / gcd(x[0], x[1])
}
