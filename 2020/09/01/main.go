package main

import (
	"bufio"
	"os"
	"strconv"
)

const preamble = 25

func main() {
	input, err := os.Open("2020/09/input.txt")
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

	for i := preamble; i < len(numbers); i++ {
		for {
			invalid, ok := isValid(numbers, i)
			if !ok {
				break
			}

			println(invalid)
			return
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
