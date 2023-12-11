package days

import (
	"AdventOfCode2023/util"
	"strings"
)

type Day11 struct{}

type galaxy struct {
	x, y int
}

func (d *Day11) Part1() int {
	return d.calcDistances(2)
}

func (d *Day11) Part2() int {
	return d.calcDistances(1000000)
}

func (d *Day11) calcDistances(extraSize int) int {
	input := util.ReadFile("input/day11.txt")

	gals := make([]galaxy, 0)
	linesEmpty := make([]bool, len(input))
	rowsEmpty := make([]bool, len(input[0]))

	for y := range input {
		linesEmpty[y] = !strings.Contains(input[y], "#")

		for x := range input[y] {
			if input[y][x] == '#' {
				gals = append(gals, galaxy{x: x, y: y})
			}
		}
	}

	for x := range input[0] {
		empty := true
		for y := 0; y < len(input) && empty; y++ {
			empty = input[y][x] != '#'
		}
		rowsEmpty[x] = empty
	}

	total := 0
	for i := range gals {
		for j := i + 1; j < len(gals); j++ {
			length := util.Abs(gals[i].x-gals[j].x) + util.Abs(gals[i].y-gals[j].y)

			for x := min(gals[i].x, gals[j].x); x < max(gals[i].x, gals[j].x); x++ {
				if rowsEmpty[x] {
					length += extraSize - 1
				}
			}

			for y := min(gals[i].y, gals[j].y); y < max(gals[i].y, gals[j].y); y++ {
				if linesEmpty[y] {
					length += extraSize - 1
				}
			}

			//println(i+1, j+1, len)
			total += length

		}
	}

	return total
}
