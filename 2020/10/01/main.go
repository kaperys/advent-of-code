package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

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
	for k, v := range jolts[1:] {
		diffs[v-jolts[k]]++
	}

	println(diffs[1] * diffs[3])
}
