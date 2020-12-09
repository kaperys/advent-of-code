package main

import (
	"bufio"
	"os"
	"strings"
)

const target = "shiny gold"

type Bag struct {
	Name     string
	Parents  []*Bag
	Children []*Bag
}

func main() {
	input, err := os.Open("2020/07/input.txt")
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

		parent, ok := rules[name]
		if !ok {
			parent = &Bag{Name: name}
			rules[name] = parent
		}

		for _, bag := range bags {
			parts := strings.Split(bag, " ")
			bag = strings.TrimSpace(strings.Join(parts[:len(parts)-1], " "))

			if bag == "no other" {
				continue
			}

			name := bag[2:]

			child, ok := rules[name]
			if !ok {
				child = &Bag{Name: name}
				rules[name] = child
			}

			child.Parents = append(child.Parents, parent)
			parent.Children = append(parent.Children, child)
		}
	}

	containers := make(map[string]int)
	walk(rules, target, containers)

	println("the number of possible containers of", target, "bags is", len(containers))
}

func walk(rules map[string]*Bag, current string, containers map[string]int) {
	c := rules[current]

	if c.Name != target {
		containers[c.Name]++
	}

	for _, p := range c.Parents {
		walk(rules, p.Name, containers)
	}
}
