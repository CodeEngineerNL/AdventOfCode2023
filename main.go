package main

import (
	"AdventOfCode2023/days"
	"fmt"
	"reflect"
	"time"
)

func runDay(any Day, name string) int64 {
	return reflect.ValueOf(any).MethodByName(name).Call(nil)[0].Int()
}

type Day interface {
	Part1() int64
	Part2() int64
}

func main() {
	println("Hello world")

	runDays := []Day{
		&days.Day1{},
	}

	for i, day := range runDays {
		start := time.Now()
		result1 := runDay(day, "Part1")
		part1Duration := time.Since(start)

		start = time.Now()
		result2 := runDay(day, "Part2")
		part2Duration := time.Since(start)

		fmt.Printf("| Day %d | %20d | %8.2f ms | %20d | %8.2f ms |",
			i+1, result1, float64(part1Duration.Microseconds())/1000, result2, float64(part2Duration.Microseconds())/1000)
	}

}
