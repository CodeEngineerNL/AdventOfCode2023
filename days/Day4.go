package days

import (
	"AdventOfCode2023/util"
	"slices"
	"strings"
)

type Day4 struct{}

func (d *Day4) Part1() int {
	input := util.ReadFile("input/day4.txt")

	total := 0

	for _, line := range input {
		cardNumbers := strings.Split(line, ":")
		numbers := strings.Split(cardNumbers[1], "|")

		winningNumbers := strings.Fields(numbers[0])
		ownNumbers := strings.Fields(numbers[1])

		cardTotal := 0

		for _, number := range ownNumbers {
			if slices.Contains(winningNumbers, number) {
				if cardTotal == 0 {
					cardTotal = 1
				} else {
					cardTotal *= 2
				}
			}
		}

		total += cardTotal
	}

	return total
}

func (d *Day4) Part2() int {
	input := util.ReadFile("input/day4.txt")

	cardResults := make([]int, len(input))
	cardNums := make([]int, len(input))

	for i, line := range input {
		cardNumbers := strings.Split(line, ":")
		numbers := strings.Split(cardNumbers[1], "|")

		winningNumbers := strings.Fields(numbers[0])
		ownNumbers := strings.Fields(numbers[1])

		cardTotal := 0

		for _, number := range ownNumbers {
			if slices.Contains(winningNumbers, number) {
				cardTotal++
			}
		}

		cardResults[i] = cardTotal
		cardNums[i] = 1
	}

	for i, numWins := range cardResults {
		if numWins > 0 {
			for cn := i + 1; cn <= i+numWins; cn++ {
				cardNums[cn] = cardNums[cn] + (1 * cardNums[i])
			}
		}
	}

	total := 0
	for _, val := range cardNums {
		total += val
	}

	return total
}
