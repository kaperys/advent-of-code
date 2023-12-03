package main

import (
	"bufio"
	"os"
	"strconv"
)

var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	input, err := os.Open("2023/01/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var total int64

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()

		var (
			li, ri int
			l, r   string
		)

		ri = len(text) - 1
		for {
			for _, v := range nums {
				if string(text[li]) == v && l == "" {
					l = v
				}

				if string(text[ri]) == v && r == "" {
					r = v
				}
			}

			if l != "" && r != "" {
				num, _ := strconv.ParseInt(l+r, 10, 64)
				total += num
				break
			}

			li++
			ri--
		}
	}

	println(total)
}
