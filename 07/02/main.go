package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const target = "shiny gold"

type Bag struct {
	Name     string
	Children []Child
}

type Child struct {
	Name   string
	Number int
}

func main() {
	input, err := os.Open("07/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	rules := make(map[string]*Bag)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "contain")

		name := strings.TrimSpace(strings.TrimSuffix(parts[0], "bags "))
		bags := strings.Split(parts[1], ",")

		parent := &Bag{Name: name}
		rules[name] = parent

		for _, bag := range bags {
			// TODO(kaperys) replace this hack with regex
			bag = strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(bag, "bags"), "bag"), "bag."), "bags."))

			if bag == "no other" {
				continue
			}

			number, err := strconv.Atoi(string(bag[0]))
			if err != nil {
				panic(err)
			}

			child := Child{
				Name:   bag[2:],
				Number: number,
			}

			parent.Children = append(parent.Children, child)
		}
	}

	println("the minimum number of bags required to carry a", target, "bags is", walk(rules, target)-1)
}

func walk(rules map[string]*Bag, current string) int {
	bags := 1

	for _, c := range rules[current].Children {
		bags += c.Number * walk(rules, c.Name)
	}

	return bags
}
