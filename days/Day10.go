package days

import (
	"AdventOfCode2023/util"
	"github.com/fatih/color"
	"math"
	"strings"
)

type Day10 struct {
	visited        [][]bool
	visitedBetween [][]bool

	width, height int
}

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

func (d *Day10) Part2() int {
	lines := d.readInput()

	d.height = len(lines)
	floodMap := make([][]uint8, d.height)
	d.width = len(lines[0])

	for i := range floodMap {
		floodMap[i] = make([]uint8, d.width)
	}

	d.visitedBetween = make([][]bool, d.height)
	for i := range d.visitedBetween {
		d.visitedBetween[i] = make([]bool, d.width)
	}

	for y := range lines {
		line := lines[y]
		for x := range line {
			if d.visited[y][x] {
				floodMap[y][x] = lines[y][x]
			} else {
				floodMap[y][x] = '.'
			}
		}
	}

	d.floodIt(floodMap)
	return d.checkInside(floodMap)
}

func isVertWall(x1, y1, x2, y2 int, m [][]uint8) bool {
	return (m[y1][x1] == 'L' && m[y2][x2] == 'J') ||
		(m[y1][x1] == 'L' && m[y2][x2] == '-') ||
		(m[y1][x1] == 'L' && m[y2][x2] == '7') ||
		(m[y1][x1] == '-' && m[y2][x2] == 'J') ||
		(m[y1][x1] == '-' && m[y2][x2] == '-') ||
		(m[y1][x1] == '-' && m[y2][x2] == '7') ||
		(m[y1][x1] == 'F' && m[y2][x2] == 'J') ||
		(m[y1][x1] == 'F' && m[y2][x2] == '-') ||
		(m[y1][x1] == 'F' && m[y2][x2] == '7')

}

func (d *Day10) countDots(m [][]uint8) (count int) {
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '.' {
				count++
			}
		}
	}
	return
}

func (d *Day10) checkInside(m [][]uint8) (count int) {
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '.' {
				if !d.shootRay(x, y, m) {
					count++
				}
			}
		}
	}
	return
}

func (d *Day10) shootRay(x, y int, m [][]uint8) (outside bool) {
	wallCountL := 0
	wallCountR := 0

	for i := y - 1; i >= 0; i-- {
		if isVertWall(x-1, i, x, i, m) {
			wallCountL++
		}
		if isVertWall(x, i, x+1, i, m) {
			wallCountR++
		}
	}

	return (wallCountL%2 == 0) || (wallCountR%2 == 0)
}

func (d *Day10) printMap(m [][]uint8) {
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	betweenC := color.New(color.FgBlue)

	for y := range m {
		for x := range m[y] {
			if m[y][x] == '.' {
				red.Print(".")
				//color.Red(".")
			} else if m[y][x] == 'O' {
				green.Print("O")
			} else if d.visitedBetween[y][x] {
				betweenC.Print(string(m[y][x]))
			} else {
				print(string(m[y][x]))
			}
		}
		println()
	}

	println()
}

func (d *Day10) floodIt(floodMap [][]uint8) {
	flooded := make([][]bool, len(floodMap))
	for i := range d.visited {
		flooded[i] = make([]bool, len(floodMap[0]))
	}

	for x := range floodMap[0] {
		d.floodFill(x, 0, floodMap)
		d.floodFill(x, d.height-1, floodMap)
	}

	for y := range floodMap {
		d.floodFill(0, 0, floodMap)
		d.floodFill(d.width-1, y, floodMap)
	}
}

func (d *Day10) floodFill(x, y int, floodMap [][]uint8) {
	if d.isValidPos(x, y, d.width, d.height) && floodMap[y][x] == '.' {
		floodMap[y][x] = 'O'

		d.floodFill(x, y+1, floodMap)
		d.floodFill(x, y-1, floodMap)

		d.floodFill(x+1, y, floodMap)
		d.floodFill(x-1, y, floodMap)

		d.floodFill(x+1, y+1, floodMap)
		d.floodFill(x+1, y-1, floodMap)

		d.floodFill(x-1, y+1, floodMap)
		d.floodFill(x-1, y-1, floodMap)
	}
}

func (d *Day10) findStartingPipe(chart []string) (x, y int) {
	for y := range chart {
		for x := range chart[y] {
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

	d.visited = make([][]bool, height)
	for i := range d.visited {
		d.visited[i] = make([]bool, width)
	}

	length := 0

	currentPipe := start

	for {
		length++

		currentSymbol := string(chart[currentPipe.y][currentPipe.x])

		x := currentPipe.x
		y := currentPipe.y
		d.visited[y][x] = true

		if strings.Contains(rightConnections, currentSymbol) && d.isValidPos(x-1, y, width, height) &&
			!d.visited[y][x-1] && strings.Contains(leftConnections, string(chart[y][x-1])) {
			currentPipe.x--
			continue
		}

		if strings.Contains(leftConnections, currentSymbol) && d.isValidPos(x+1, y, width, height) &&
			!d.visited[y][x+1] && strings.Contains(rightConnections, string(chart[y][x+1])) {
			currentPipe.x++
			continue

		}

		if strings.Contains(bottomConnections, currentSymbol) && d.isValidPos(x, y-1, width, height) &&
			!d.visited[y-1][x] && strings.Contains(topConnections, string(chart[y-1][currentPipe.x])) {
			currentPipe.y--
			continue

		}

		if strings.Contains(topConnections, currentSymbol) && d.isValidPos(x, y+1, width, height) &&
			!d.visited[y+1][x] && strings.Contains(bottomConnections, string(chart[y+1][x])) {
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

func (d *Day10) readInput() []string {
	return util.ReadFile("input/day10.txt")
}
