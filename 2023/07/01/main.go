package main

import (
	"bufio"
	"os"
	"slices"

	"github.com/kaperys/advent-of-code/aoc/strings"
	"golang.org/x/exp/maps"
)

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	Type int
	Bid  int

	Cards         string
	CardsRemapped string
}

var remap = map[rune]string{
	'A': "B",
	'K': "C",
	'Q': "D",
	'J': "E",
	'T': "F",
	'9': "G",
	'8': "H",
	'7': "I",
	'6': "J",
	'5': "K",
	'4': "L",
	'3': "M",
	'2': "N",
}

func main() {
	input, err := os.Open("2023/07/input.txt")
	if err != nil {
		panic(err)
	}

	defer input.Close()

	var scores []Hand

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		h := Hand{Cards: text[:5], Bid: strings.ToInt(text[6:])}

		counts := make(map[rune]int)
		for _, card := range h.Cards {
			counts[card]++
			h.CardsRemapped += remap[card]
		}

		hands := maps.Values(counts)
		slices.Sort(hands)

		switch {
		case slices.Equal(hands, []int{1, 1, 1, 1, 1}):
			h.Type = HighCard
		case slices.Equal(hands, []int{1, 1, 1, 2}):
			h.Type = OnePair
		case slices.Equal(hands, []int{1, 2, 2}):
			h.Type = TwoPair
		case slices.Equal(hands, []int{1, 1, 3}):
			h.Type = ThreeOfAKind
		case slices.Equal(hands, []int{2, 3}):
			h.Type = FullHouse
		case slices.Equal(hands, []int{1, 4}):
			h.Type = FourOfAKind
		case slices.Equal(hands, []int{5}):
			h.Type = FiveOfAKind
		}

		scores = append(scores, h)
	}

	slices.SortFunc(scores, func(a, b Hand) int {
		switch {
		case a.Type > b.Type:
			return 1
		case a.Type < b.Type:
			return -1
		}

		if a.CardsRemapped > b.CardsRemapped {
			return 1
		}

		return -1
	})

	var winnings int

	for i := 0; i < len(scores); i++ {
		winnings += (len(scores) - i) * scores[i].Bid
	}

	println(winnings)
}
