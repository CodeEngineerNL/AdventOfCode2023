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

	println("total", total)
	return total
}

func (d *Day1) getNumberFromLine(line string) int {
	var i int
	for i = 0; i < len(line) && !d.isDigit(line[i]); i++ {
		/* empty */
	}

	first := string(line[i])

	for i = len(line) - 1; i > 0 && !d.isDigit(line[i]); i-- {
		/* empty */
	}

	last := string(line[i])
	digit, _ := strconv.Atoi(first + last)
	return digit
}

func (d *Day1) isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func (d *Day1) scanIt(line string) int {
	num := d.scanForDigit(line)
	num2 := d.scanForLastDigit(line)

	return num*10 + num2
}

func (d *Day1) scanForDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if d.isDigit(line[i]) {
			return int(line[i] - '0')
		}

		for scanI, val := range toScan {
			if strings.HasPrefix(line[i:], val) {
				return scanI
			}
		}
	}

	return 0
}

func (d *Day1) scanForLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if d.isDigit(line[i]) {
			return int(line[i] - '0')
		}

		for scanI, val := range toScan {
			if strings.HasPrefix(line[i:], val) {
				return scanI
			}
		}
	}

	return 0
}
