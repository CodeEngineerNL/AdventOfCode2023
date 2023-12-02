package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d *Day2) Part1() int {
	input := d.getInput()

	total := 0
	for i, line := range input {
		if d.isValidLine(line) {
			total += i + 1
		}
	}

	return total
}

func (d *Day2) Part2() int {
	input := d.getInput()

	total := 0
	for _, line := range input {
		total += d.getLinePower(line)
	}

	return total
}

func (d *Day2) isValidLine(line string) bool {
	gameSplit := strings.Split(line, ":")
	hands := strings.Split(gameSplit[1], ";")

	for _, hand := range hands {
		cubes := strings.Split(hand, ",")
		for _, set := range cubes {
			set := strings.TrimSpace(set)
			setData := strings.Split(set, " ")

			num, _ := strconv.Atoi(setData[0])
			color := setData[1]

			switch color {
			case "red":
				if num > 12 {
					return false
				}
			case "green":
				if num > 13 {
					return false
				}
			case "blue":
				if num > 14 {
					return false
				}
			}
		}
	}

	return true
}

func (d *Day2) getLinePower(line string) int {
	gameSplit := strings.Split(line, ":")
	hands := strings.Split(gameSplit[1], ";")

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, hand := range hands {
		cubes := strings.Split(hand, ",")
		for _, set := range cubes {
			set := strings.TrimSpace(set)
			setData := strings.Split(set, " ")

			num, _ := strconv.Atoi(setData[0])
			color := setData[1]

			switch color {
			case "red":
				if num > maxRed {
					maxRed = num
				}
			case "green":
				if num > maxGreen {
					maxGreen = num
				}
			case "blue":
				if num > maxBlue {
					maxBlue = num
				}
			}
		}
	}

	return maxRed * maxGreen * maxBlue
}

func (d *Day2) getInput() []string {
	f, err := os.Open("input/day2.txt")
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
