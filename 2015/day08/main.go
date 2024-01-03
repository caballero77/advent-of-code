package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	count := part1(input)
	fmt.Printf("Part 1: %d\n", count)

	count = part2(input)
	fmt.Printf("Part 2: %d\n", count)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += len(line) - getCharactersCount(line)
	}
	return sum
}

func getCharactersCount(line string) int {
	length := 0
	for i := 1; i < len(line)-1; i++ {
		if line[i] == '\\' {
			if line[i] == line[i+1] {
				i++
			} else if i < len(line)-1 && line[i+1] == '"' {
				i++
			} else if i+3 < len(line)-1 {
				if ok, _ := regexp.MatchString("[^0-9A-Fa-f]", line[i:i+3]); ok {
					i += 3
				}
			}

		}
		length++
	}
	return length
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += len(fmt.Sprintf("%+q", line)) - len(line)
	}
	return sum
}
