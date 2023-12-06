package days

import (
	"AdventOfCode2023/util"
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct{}

func (d *Day5) Part1() int {
	seeds, maps := d.readInput()

	for i := range maps {
		slices.SortFunc(maps[i], func(a, b entry) int {
			return cmp.Compare(a.source, b.source)
		})
	}

	lowestDest := math.MaxInt

	for _, seed := range seeds {
		newSeed := seed
		for i, _ := range maps {
			newSeed = d.findDest(newSeed, maps[i])
		}
		lowestDest = min(lowestDest, newSeed)
	}

	return lowestDest
}

func (d *Day5) findDest(num int, mapEntries []entry) int {
	found := false
	for mapI := range mapEntries {
		sourceEnd := mapEntries[mapI].source + mapEntries[mapI].num
		found = num >= mapEntries[mapI].source && num < sourceEnd
		if found {
			return mapEntries[mapI].dest + (num - mapEntries[mapI].source)
		}
	}

	return num
}

func (d *Day5) Part2() int {
	seeds, maps := d.readInput()

	seedEntries := make([]entry, 0)
	for i := 0; i < len(seeds); i += 2 {
		seedEntries = append(seedEntries, entry{
			source: 0,
			dest:   seeds[i],
			num:    seeds[i+1],
		})
	}

	for i := range maps {
		slices.SortFunc(maps[i], func(a, b entry) int {
			return cmp.Compare(a.dest, b.dest)
		})
	}

	slices.SortFunc(seedEntries, func(a, b entry) int {
		return cmp.Compare(a.dest, b.dest)
	})

	for i := 0; ; i++ {

		source := i
		for mapI := len(maps) - 1; mapI >= 0; mapI-- {
			source = d.findSource(source, maps[mapI])
			if source < 0 {
				break
			}
		}

		if source < 0 {
			break
		}

		for seedI := range seedEntries {
			currentRange := seedEntries[seedI]

			if source >= currentRange.dest && source < currentRange.dest+currentRange.num {
				return i
			}
		}
	}

	return math.MaxInt
}

func (d *Day5) findSource(num int, mapEntries []entry) int {
	found := false
	for mapI := range mapEntries {
		destEnd := mapEntries[mapI].dest + mapEntries[mapI].num
		found = num >= mapEntries[mapI].dest && num < destEnd
		if found {
			return num - (mapEntries[mapI].dest - mapEntries[mapI].source)
		}
	}

	return num
}

type entry struct {
	source int
	dest   int
	num    int
}

func (d *Day5) readInput() ([]int, [][]entry) {
	input := util.ReadFile("input/day5.txt")

	seedsString := strings.Fields(strings.Split(input[0], ":")[1])

	seeds := make([]int, 0, len(seedsString))

	for _, seed := range seedsString {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}

	i := 1
	for ; input[i] != "seed-to-soil map:"; i++ {
		// empty
	}
	i++
	maps := make([][]entry, 0)
	mapEntries := make([]entry, 0)

	for ; i < len(input); i++ {
		if input[i] == "" {
			maps = append(maps, mapEntries)
			mapEntries = make([]entry, 0)
			i++
			continue
		}

		entryNums := strings.Fields(input[i])

		dest, _ := strconv.Atoi(entryNums[0])
		source, _ := strconv.Atoi(entryNums[1])
		num, _ := strconv.Atoi(entryNums[2])

		mapEntries = append(mapEntries, entry{
			source: source,
			dest:   dest,
			num:    num,
		})
	}

	maps = append(maps, mapEntries)

	return seeds, maps
}
