package days

import (
	"AdventOfCode2023/util"
	"math"
	"strings"
)

type Day10 struct{}

type piece struct {
	x, y int
	//up, down, left, right *piece
}

const (
	topConnections    = "|7FS"
	bottomConnections = "|LJS"
	leftConnections   = "-LFS"
	rightConnections  = "-J7S"
)

func (d *Day10) Part1() int {
	lines := d.readInput()

	x, y := d.findStartingPipe(lines)
	pipeLen := d.parseLoopLength(x, y, lines)

	return int(math.Round(float64(pipeLen) / 2))
}

func (d *Day10) findStartingPipe(chart []string) (x, y int) {
	for y := 0; y < len(chart); y++ {
		for x := 0; x < len(chart[y]); x++ {
			if chart[y][x] == 'S' {
				return x, y
			}
		}
	}

	return -1, -1
}

func (d *Day10) parseLoopLength(sx, sy int, chart []string) int {
	width := len(chart[0])
	height := len(chart)

	start := &piece{
		x: sx,
		y: sy,
	}

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	length := 0

	currentPipe := start

	for {
		length++

		currentSymbol := string(chart[currentPipe.y][currentPipe.x])

		x := currentPipe.x
		y := currentPipe.y
		visited[y][x] = true

		if strings.Contains(rightConnections, currentSymbol) && d.isValidPos(x-1, y, width, height) &&
			!visited[y][x-1] && strings.Contains(leftConnections, string(chart[y][x-1])) {
			currentPipe.x--
			continue
		}

		if strings.Contains(leftConnections, currentSymbol) && d.isValidPos(x+1, y, width, height) &&
			!visited[y][x+1] && strings.Contains(rightConnections, string(chart[y][x+1])) {
			currentPipe.x++
			continue

		}

		if strings.Contains(bottomConnections, currentSymbol) && d.isValidPos(x, y-1, width, height) &&
			!visited[y-1][x] && strings.Contains(topConnections, string(chart[y-1][currentPipe.x])) {
			currentPipe.y--
			continue

		}

		if strings.Contains(topConnections, currentSymbol) && d.isValidPos(x, y+1, width, height) &&
			!visited[y+1][x] && strings.Contains(bottomConnections, string(chart[y+1][x])) {
			currentPipe.y++
			continue
		}

		if currentPipe.x == start.x && currentPipe.y == start.y {
			return length
		}
	}
}

func (d *Day10) isValidPos(x, y, width, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}

func (d *Day10) Part2() int {
	lines := d.readInput()
	return len(lines)
}

func (d *Day10) readInput() []string {
	return util.ReadFile("input/day10.txt")
}
