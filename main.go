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
		&days.Day2{},
	}

	for i, day := range runDays {
		start1 := time.Now()
		result1 := runDay(day, "Part1")
		part1Duration := time.Since(start1)

		start2 := time.Now()
		result2 := runDay(day, "Part2")
		part2Duration := time.Since(start2)

		fmt.Printf("| Day %d | %20d | %10s | %20d | %10s |\r\n", i+1, result1, part1Duration, result2, part2Duration)
	}

}
