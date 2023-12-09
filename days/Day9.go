package days

import (
	"AdventOfCode2023/util"
	"strconv"
	"strings"
)

type Day9 struct{}

func (d *Day9) Part1() int {
	lines := d.readInput()

	total := 0
	for l := range lines {
		line := lines[l]
		lastDiffs := getLastDiffs(line, false)
		predictions := make([]int, len(lastDiffs))

		for i := len(lastDiffs) - 1; i > 0; i-- {
			predictions[i-1] = predictions[i] + lastDiffs[i-1]
		}

		prediction := predictions[0] + line[len(line)-1]
		total += prediction
	}

	return total
}

func getLastDiffs(nums []int, history bool) []int {
	diffs := make([]int, len(nums))
	copy(diffs, nums)

	numDiffs := len(diffs) - 1
	numZeros := 0

	lastDiffs := make([]int, 0)

	for {
		numZeros = 0
		for i := 0; i < numDiffs; i++ {
			diffs[i] = diffs[i+1] - diffs[i]
			if diffs[i] == 0 {
				numZeros++
			}
		}

		if history {
			lastDiffs = append(lastDiffs, diffs[0])
		} else {
			lastDiffs = append(lastDiffs, diffs[numDiffs-1])
		}

		if numDiffs == numZeros {
			break
		}

		numDiffs--
	}

	return lastDiffs
}

func (d *Day9) Part2() int {
	lines := d.readInput()

	total := 0
	for l := range lines {
		line := lines[l]
		lastDiffs := getLastDiffs(line, true)
		predictions := make([]int, len(lastDiffs))

		for i := len(lastDiffs) - 1; i > 0; i-- {
			predictions[i-1] = lastDiffs[i-1] - predictions[i]
		}

		prediction := line[0] - predictions[0]
		total += prediction
	}

	return total
}

func (d *Day9) readInput() (values [][]int) {
	lines := util.ReadFile("input/day9.txt")

	for i := range lines {
		numsEntry := make([]int, 0)

		numsStr := strings.Fields(lines[i])

		for j := range numsStr {
			num, _ := strconv.Atoi(numsStr[j])
			numsEntry = append(numsEntry, num)
		}

		values = append(values, numsEntry)
	}

	return
}
