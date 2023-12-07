package main

import (
	"bufio"
	"os"
	"slices"
	"strings"

	ustrings "github.com/kaperys/advent-of-code/aoc/strings"
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
	'T': "E",
	'9': "F",
	'8': "G",
	'7': "H",
	'6': "I",
	'5': "J",
	'4': "K",
	'3': "L",
	'2': "M",
	'J': "N",
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
		h := Hand{Cards: text[:5], Bid: ustrings.ToInt(text[6:])}

		counts := make(map[rune]int)
		for _, card := range h.Cards {
			counts[card]++
			h.CardsRemapped += remap[card]
		}

		freq := mostFrequentCard(counts)
		for _, card := range h.Cards {
			if card == 'J' {
				counts[card]--
				counts[freq]++

				if counts[card] == 0 {
					delete(counts, card)
				}
			}
		}
		h.Cards = strings.Replace(h.Cards, "J", string(freq), -1)

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

func mostFrequentCard(cards map[rune]int) rune {
	var (
		card rune
		num  int
	)

	for c, n := range cards {
		if n > num && c != 'J' {
			card = c
			num = n
		}
	}

	if card == 'J' {
		// if the most frequent card is a joker,
		// replace it with the highest value card
		return 'A'
	}

	return card
}
