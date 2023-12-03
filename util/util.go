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
