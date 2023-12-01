package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]string{
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

// It looks like some of the digits are actually spelled out with letters: one, two,
// three, four, five, six, seven, eight, and nine also count as valid "digits"
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

		var li, ri int
		ri = len(text) - 1
		var lv, rv, l, r string

		for {
			lv += string(text[li])
			rv = string(text[ri]) + rv

			for k, v := range nums {
				if (strings.HasSuffix(lv, v) || strings.HasSuffix(lv, k)) && l == "" {
					l = k
				}

				if (strings.HasPrefix(rv, v) || strings.HasPrefix(rv, k)) && r == "" {
					r = k
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
