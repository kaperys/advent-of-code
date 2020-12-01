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

	var expenses []float64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		expense, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}

		expenses = append(expenses, expense)
	}

	for _, x := range expenses {
		for _, y := range expenses {
			if x+y == target {
				fmt.Printf("%v x %v = %f\n", x, y, x*y)
				return
			}
		}
	}
}
