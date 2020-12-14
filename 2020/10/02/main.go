package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// https://adventofcode.com/2020/day/10#part2
func main() {
	input, err := os.Open("2020/10/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var jolts []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		jolts = append(jolts, i)
	}

	sort.Ints(jolts)
	jolts = append([]int{0}, append(jolts, jolts[len(jolts)-1]+3)...)

	diffs := make(map[int]int)

	combinations := make(map[int]int)
	combinations[0]++

	for k, v := range jolts[1:] {
		diffs[v-jolts[k]]++
		combinations[v] = combinations[v-1] + combinations[v-2] + combinations[v-3]
	}

	println(combinations[jolts[len(jolts)-1]])
}
