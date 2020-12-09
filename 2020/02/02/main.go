package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("2020/02/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var validPasswords int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		policy := strings.Split(parts[0], " ")
		limits := strings.Split(policy[0], "-")

		first, _ := strconv.Atoi(limits[0])
		second, _ := strconv.Atoi(limits[1])
		letter := policy[1]
		password := strings.TrimSpace(parts[1])

		var occurances int
		if string(password[first-1]) == letter {
			occurances++
		}

		if string(password[second-1]) == letter {
			occurances++
		}

		if occurances == 1 {
			validPasswords++
		}
	}

	println("there are", validPasswords, "valid passwords")
}
