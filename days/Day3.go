package days

import (
	"AdventOfCode2023/util"
	"strconv"
)

type Day3 struct{}

func (d *Day3) Part1() int {
	input := util.ReadFile("input/Day3.txt")

	total := 0

	for y := 0; y < len(input); y++ {

		numStart := -1
		numEnd := -1
		// Add an extra dot for the algorithm to work when a number is on the end of a line
		line := input[y] + "."
		for x := 0; x < len(line); x++ {
			if !util.IsDigit(line[x]) {
				if numStart != -1 {
					numEnd = x - 1

					num, _ := strconv.Atoi(line[numStart : numEnd+1])

					if d.scanForSymbol(numStart, numEnd, y, input) {
						total += num
					}

					numStart = -1
					numEnd = -1
				}

			} else {
				if numStart == -1 {
					numStart = x
				}
			}
		}
	}

	return total
}

func (d *Day3) scanForSymbol(startx, endx int, yPos int, symbolMap []string) bool {
	width := len(symbolMap[0])
	height := len(symbolMap)

	for y := yPos - 1; y <= yPos+1; y++ {
		for x := startx - 1; x <= endx+1; x++ {
			if d.isValidPosition(x, y, width, height) && !util.IsDigit(symbolMap[y][x]) && symbolMap[y][x] != '.' {
				return true
			}
		}
	}

	return false
}

func (d *Day3) isValidPosition(x, y, width, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}

func (d *Day3) Part2() int {
	return 0
}
