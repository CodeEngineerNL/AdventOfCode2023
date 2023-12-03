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

func (d *Day3) Part2() int {
	input := util.ReadFile("input/Day3.txt")

	// Add a symbol before and after each line for the algorithm to work
	for y, _ := range input {
		input[y] = "." + input[y] + "."
	}

	total := 0

	for y := 0; y < len(input); y++ {
		line := input[y]

		for x := 0; x < len(line); x++ {
			if line[x] == '*' {
				num1, num2, found := d.scanForNumbers(x, y, input)
				if found {
					total += num1 * num2
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

func (d *Day3) scanForNumbers(posx, posy int, symbolMap []string) (int, int, bool) {
	width := len(symbolMap[0])
	height := len(symbolMap)

	var nums []int

	for y := posy - 1; y <= posy+1; y++ {
		line := symbolMap[y]
		for x := posx - 1; x <= posx+1; x++ {
			if d.isValidPosition(x, y, width, height) && util.IsDigit(symbolMap[y][x]) {
				startPos := x
				endPos := x

				for startPos = x; util.IsDigit(line[startPos]); startPos-- {
					// empty
				}

				for endPos = x; util.IsDigit(line[endPos]); endPos++ {
					// empty
				}

				// We scanned to endPos. Change x to prevent scanning this position again or else me might get the same number again
				x = endPos

				if startPos >= 0 && endPos-1 < len(line) {
					foundNum, _ := strconv.Atoi(line[startPos+1 : endPos])
					nums = append(nums, foundNum)

					if len(nums) == 2 {
						return nums[0], nums[1], true
					}
				}
			}
		}
	}

	return 0, 0, false
}

func (d *Day3) isValidPosition(x, y, width, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}
