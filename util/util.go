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
