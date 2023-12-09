package strings

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ToIntSlice(s string) []int {
	var i []int

	for _, v := range strings.Fields(s) {
		i = append(i, ToInt(v))
	}

	return i
}
