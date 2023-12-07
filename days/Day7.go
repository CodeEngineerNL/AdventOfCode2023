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

	order  = "23456789TJQKA"
	order2 = "J23456789TQKA"
)

type move struct {
	hand     string
	bet      int
	handType HandType
}

func (d *Day7) Part1() int {
	lines := d.readInput(d.getHandType)

	slices.SortFunc(lines, func(a, b move) int {
		sort := cmp.Compare(a.handType, b.handType)
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
	lines := d.readInput(d.getHandTypeWithJokers)

	slices.SortFunc(lines, func(a, b move) int {
		sort := cmp.Compare(a.handType, b.handType)
		if sort != 0 {
			return sort
		}

		for i := range a.hand {
			numA := strings.Index(order2, string(a.hand[i]))
			numB := strings.Index(order2, string(b.hand[i]))

			sort = cmp.Compare(numA, numB)
			if sort != 0 {
				return sort
			}
		}

		return sort
	})

	for _, line := range lines {
		println(line.hand, line.handType)
	}

	score := 0
	for i := range lines {
		score = score + (i+1)*lines[i].bet
	}

	return score
}

func (d *Day7) getHandTypeWithJokers(hand string) HandType {
	cards := make(map[int32]int)
	for _, c := range hand {
		cards[c] = cards[c] + 1
	}

	pairs := 0
	threes := 0

	numJokers := cards['J']

	for _, v := range cards {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			if numJokers == 1 {
				return FiveOfAKind
			} else if numJokers == 4 {
				return FiveOfAKind
			}
			return FourOfAKind
		case 3:
			threes++
		case 2:
			pairs++
		}
	}

	if pairs == 1 && threes == 1 {
		if numJokers == 2 || numJokers == 3 {
			return FiveOfAKind
		}

		return FullHouse
	}

	if threes == 1 {
		if numJokers == 1 {
			return FourOfAKind
		} else if numJokers == 3 {
			return FourOfAKind
		}
		return ThreeOfAKind
	}

	if pairs == 2 {
		if numJokers == 1 {
			return FullHouse
		} else if numJokers == 2 {
			return FourOfAKind
		}

		return TwoPair
	}
	if pairs == 1 {
		if numJokers == 1 {
			return ThreeOfAKind
		} else if numJokers == 2 {
			return ThreeOfAKind
		}

		return OnePair
	}

	if numJokers == 1 {
		return OnePair
	}

	return HighCard
}

func (d *Day7) readInput(getHandType func(string) HandType) []move {
	lines := util.ReadFile("input/day7.txt")
	hands := make([]move, 0, len(lines))

	for i := range lines {
		parts := strings.Fields(lines[i])
		bet, _ := strconv.Atoi(parts[1])
		hands = append(hands, move{
			hand:     parts[0],
			bet:      bet,
			handType: getHandType(parts[0]),
		})
	}

	return hands
}
