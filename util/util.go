package util

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Could not read file ", filename)
	}

	result := strings.Split(string(input), "\r\n")

	return result
}

func IsDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func FindLCM(nums []int) int {
	lcm := 1
	divisor := 2

	for true {
		counter := 0
		divisible := false

		for i := 0; i < len(nums); i++ {
			if nums[i] == 1 {
				counter++
			}

			if nums[i]%divisor == 0 {
				divisible = true
				nums[i] = nums[i] / divisor
			}
		}

		if divisible {
			lcm = lcm * divisor
		} else {
			divisor++
		}

		if counter == len(nums) {
			return lcm
		}
	}

	return 0
}
