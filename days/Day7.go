package days

import (
	"AdventOfCode2023/util"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Day7 struct{}

type HandType int

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind

	order = "23456789TJQKA"
)

type move struct {
	hand     string
	bet      int
	handType HandType
}

func (d *Day7) Part1() int {
	lines := d.readInput()

	slices.SortFunc(lines, func(a, b move) int {
		sort := cmp.Compare(d.getHandType(a.hand), d.getHandType(b.hand))
		if sort != 0 {
			return sort
		}

		for i := range a.hand {
			numA := strings.Index(order, string(a.hand[i]))
			numB := strings.Index(order, string(b.hand[i]))

			sort = cmp.Compare(numA, numB)
			if sort != 0 {
				return sort
			}
		}

		return sort
	})

	score := 0
	for i := range lines {
		score = score + (i+1)*lines[i].bet
	}

	return score
}

func (d *Day7) getHandType(hand string) HandType {
	cards := make(map[int32]int)
	for _, c := range hand {
		cards[c] = cards[c] + 1
	}

	pairs := 0
	threes := 0

	for _, v := range cards {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			threes++
		case 2:
			pairs++
		}
	}

	if pairs == 1 && threes == 1 {
		return FullHouse
	}
	if threes == 1 {
		return ThreeOfAKind
	}
	if pairs == 2 {
		return TwoPair
	}
	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func (d *Day7) Part2() int {
	lines := d.readInput()

	return len(lines)
}

func (d *Day7) readInput() []move {
	lines := util.ReadFile("input/day7.txt")
	hands := make([]move, 0, len(lines))

	for i := range lines {
		parts := strings.Fields(lines[i])
		bet, _ := strconv.Atoi(parts[1])
		hands = append(hands, move{
			hand:     parts[0],
			bet:      bet,
			handType: d.getHandType(parts[0]),
		})
	}

	return hands
}
