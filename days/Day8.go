package days

import (
	"AdventOfCode2023/util"
	"strings"
)

type Day8 struct{}

type mapItem struct {
	name    string
	left    *mapItem
	right   *mapItem
	visited bool
}

func (d *Day8) Part1() int {
	moves, items := d.readInput()

	var position *mapItem
	for i := range items {
		if items[i].name == "AAA" {
			position = items[i]
			break
		}
	}

	if position == nil {
		return 0
	}

	steps := 0
	for true {
		for _, m := range moves {
			steps++

			if m == 'L' {
				position = position.left
			} else if m == 'R' {
				position = position.right
			}

			if position.name == "ZZZ" {
				return steps
			}
		}
	}

	return 0
}

func (d *Day8) Part2() int {
	moves, items := d.readInput()

	positions := make([]*mapItem, 0)

	for i := range items {
		if items[i].name[len(items[i].name)-1] == 'A' {
			positions = append(positions, items[i])
		}
	}

	moveCounters := make([]int, len(positions))

	for i := range positions {
		moveCounters[i] = d.getStepsToZ(moves, positions[i])
	}

	return d.findLCM(moveCounters)
}

func (d *Day8) findLCM(nums []int) int {
	powered := make([]int, len(nums))

	copy(powered, nums)

	for {
		i := 0
		for ; i < len(powered)-1; i++ {
			if powered[i] < powered[i+1] {
				break
			}
		}

		powered[i] += nums[i]

		allEqual := true
		for i := 1; i < len(powered) && allEqual; i++ {
			if powered[i-1] != powered[i] {
				allEqual = false
			}
		}
		if allEqual {
			break
		}
	}

	return powered[0]
}

func (d *Day8) getStepsToZ(moves string, item *mapItem) int {
	steps := 0
	currentItem := item

	for currentItem.name[len(currentItem.name)-1] != 'Z' {
		move := moves[steps%len(moves)]
		if move == 'L' {
			currentItem = currentItem.left
		} else {
			currentItem = currentItem.right
		}
		steps++
	}

	return steps
}

func (d *Day8) readInput() (string, []*mapItem) {
	lines := util.ReadFile("input/day8.txt")
	moves := lines[0]

	existingItems := make(map[string]*mapItem)

	for i := 2; i < len(lines); i++ {
		line := strings.ReplaceAll(lines[i], "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")

		fields := strings.Fields(line)

		name := fields[0]
		leftName := fields[2]
		rightName := fields[3]

		leftItem := &mapItem{name: leftName}
		rightItem := &mapItem{name: rightName}

		item := &mapItem{
			name:    name,
			left:    leftItem,
			right:   rightItem,
			visited: false,
		}

		existingItems[name] = item
	}

	items := make([]*mapItem, 0)
	for key := range existingItems {
		currentItem := existingItems[key]

		currentItem.left = existingItems[currentItem.left.name]
		currentItem.right = existingItems[currentItem.right.name]
		items = append(items, currentItem)
	}

	return moves, items
}
