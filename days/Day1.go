package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day1 struct{}

type numStrPos struct {
	str string
	pos int
}

func (d *Day1) Part1() int64 {
	input := d.getInput()

	total := int64(0)
	for _, line := range input {
		digit := d.getNumberFromLine(line)
		total += int64(digit)
	}

	return total
}

func (d *Day1) Part2() int64 {
	input := d.getInput()

	total := int64(0)
	for _, line := range input {
		res := int64(d.scanIt(line))
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
	toScan := []numStrPos{
		{"zero", 0},
		{"one", 0},
		{"two", 0},
		{"three", 0},
		{"four", 0},
		{"five", 0},
		{"six", 0},
		{"seven", 0},
		{"eight", 0},
		{"nine", 0},
	}

	num := d.scanForDigit(line, toScan)
	num2 := d.scanForLastDigit(line, toScan)

	return num*10 + num2
}

func (d *Day1) scanForDigit(line string, toScan []numStrPos) int {
	for i := 0; i < len(line); i++ {
		if d.isDigit(line[i]) {
			return int(line[i] - '0')
		}

		for val := range toScan {
			scanThis := &toScan[val]

			if scanThis.str[scanThis.pos] == line[i] {
				scanThis.pos = scanThis.pos + 1
			} else if scanThis.str[0] == line[i] {
				scanThis.pos = 1
			} else {
				scanThis.pos = 0
			}

			if scanThis.pos == len(scanThis.str) {
				return val
			}
		}
	}

	return 0
}

func (d *Day1) scanForLastDigit(line string, toScan []numStrPos) int {
	for i := range toScan {
		toScan[i].pos = len(toScan[i].str) - 1
	}

	for i := len(line) - 1; i >= 0; i-- {
		if d.isDigit(line[i]) {
			return int(line[i] - '0')
		}

		for val := range toScan {
			scanThis := &toScan[val]

			if scanThis.str[scanThis.pos] == line[i] {
				scanThis.pos = scanThis.pos - 1
			} else if scanThis.str[len(scanThis.str)-1] == line[i] {
				scanThis.pos = len(scanThis.str) - 2
			} else {
				scanThis.pos = len(scanThis.str) - 1
			}

			if scanThis.pos == -1 {
				return val
			}
		}
	}

	return 0
}

func (d *Day1) getInput() []string {
	f, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal("Could not open input file")
	}
	defer f.Close()

	var result []string

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
