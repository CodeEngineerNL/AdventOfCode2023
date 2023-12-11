package days

import (
	"AdventOfCode2023/util"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d *Day1) Part1() int {
	input := util.ReadFile("input/Day1.txt")

	total := 0
	for _, line := range input {
		digit := d.getNumberFromLine(line)
		total += digit
	}

	return total
}

var toScan = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func (d *Day1) Part2() int {
	input := util.ReadFile("input/Day1.txt")

	total := 0
	for _, line := range input {
		res := d.scanIt(line)
		total += res
	}

	return total
}

func (d *Day1) getNumberFromLine(line string) int {
	var i int
	for i = 0; i < len(line) && !util.IsDigit(line[i]); i++ {
		/* empty */
	}

	first := string(line[i])

	for i = len(line) - 1; i > 0 && !util.IsDigit(line[i]); i-- {
		/* empty */
	}

	last := string(line[i])
	digit, _ := strconv.Atoi(first + last)
	return digit
}

func (d *Day1) scanIt(line string) int {
	num := 0
	done := false
	for i := 0; i < len(line) && !done; i++ {
		num, done = d.getDigitOnPosition(line, i)
	}

	num2 := 0
	done = false
	for i := len(line) - 1; i >= 0 && !done; i-- {
		num2, done = d.getDigitOnPosition(line, i)
	}

	return num*10 + num2
}

func (d *Day1) getDigitOnPosition(line string, i int) (int, bool) {
	if util.IsDigit(line[i]) {
		return int(line[i] - '0'), true
	}

	for scanI, val := range toScan {
		if strings.HasPrefix(line[i:], val) {
			return scanI, true
		}
	}
	return 0, false
}
