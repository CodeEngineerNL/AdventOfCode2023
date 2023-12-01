package main

import (
	"AdventOfCode2023/days"
	"fmt"
	"reflect"
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
		result1 := runDay(day, "Part1")
		result2 := runDay(day, "Part2")

		fmt.Printf("| Day %d | %20d | %20d |", i, result1, result2)
	}

}
