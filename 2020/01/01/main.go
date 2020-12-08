package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const target = 2020

func main() {
	input, err := os.Open("01/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	expenses := make(map[float64]struct{})

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		expense, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		expenses[expense] = struct{}{}
	}

	for x := range expenses {
		y := target - x
		_, ok := expenses[y]
		if ok {
			fmt.Printf("%0.f\n", x*y)
			return
		}
	}
}
