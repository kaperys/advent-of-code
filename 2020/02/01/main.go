package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("02/input.txt")
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

		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])
		letter := policy[1]
		password := strings.TrimSpace(parts[1])

		var appearances int
		for _, char := range password {
			if string(char) == letter {
				appearances++
			}
		}

		if appearances >= min && appearances <= max {
			validPasswords++
		}
	}

	println("there are", validPasswords, "valid passwords")
}
