package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("2020/12/input.txt")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(strings.TrimSpace(string(input)), "\n")
	_ = instructions
}
