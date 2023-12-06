package days

import (
	"AdventOfCode2023/util"
	"strconv"
	"strings"
)

type Day6 struct{}

func (d *Day6) Part1() int {
	lines := d.readInput()

	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	total := 0
	for i := range times {

		time, _ := strconv.Atoi(times[i])
		recordDistance, _ := strconv.Atoi(distances[i])

		raceTotal := d.calcRaceTOtal(time, recordDistance)

		if total != 0 {
			total *= raceTotal
		} else {
			total = raceTotal
		}
	}

	return total
}

func (d *Day6) calcRaceTOtal(time int, recordDistance int) int {
	raceTotal := 0

	for holdTime := 0; holdTime <= time; holdTime++ {
		travelTime := time - holdTime
		distance := holdTime * travelTime

		if distance > recordDistance {
			raceTotal++
		}
	}
	return raceTotal
}

func (d *Day6) Part2() int {
	lines := d.readInput()

	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))

	return d.calcRaceTOtal(time, distance)
}

func (d *Day6) readInput() []string {
	return util.ReadFile("input/day6.txt")
}
