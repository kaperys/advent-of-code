package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const preamble = 25

func main() {
	input, err := os.Open("09/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var numbers []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, i)
	}

	var xmas int
	for i := preamble; i < len(numbers); i++ {
		for {
			invalid, ok := isValid(numbers, i)
			if !ok {
				break
			}

			xmas = invalid
			break
		}
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			var sum int
			for _, n := range numbers[i : j+1] {
				sum += n

				if sum == xmas {
					sort.Ints(numbers[i : j+1])

					println(numbers[i] + numbers[j])
					return
				}
			}
		}
	}
}

func isValid(numbers []int, i int) (int, bool) {
	for j := i - preamble; j < i; j++ {
		for k := j + 1; k < i; k++ {
			if numbers[j]+numbers[k] == numbers[i] {
				return 0, false
			}
		}
	}

	return numbers[i], true
}
