package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	count := part1(input)
	fmt.Printf("Part 1: %d\n", count)

	count = part2(input)
	fmt.Printf("Part 1: %d\n", count)
}

func part2(input string) int {
	for i := 0; i < 50; i++ {
		curr := strings.Builder{}
		prev := input[0]
		count := 1
		for j := 1; j < len(input); j++ {
			if input[j] == prev {
				count++
				continue
			} else {
				curr.WriteString(strconv.Itoa(count) + string(prev))
				prev = input[j]
				count = 1
			}
		}
		curr.WriteString(strconv.Itoa(count) + string(prev))
		input = curr.String()
	}
	return len(input)
}

func part1(input string) int {
	for i := 0; i < 40; i++ {
		curr := strings.Builder{}
		prev := input[0]
		count := 1
		for j := 1; j < len(input); j++ {
			if input[j] == prev {
				count++
				continue
			} else {
				curr.WriteString(strconv.Itoa(count) + string(prev))
				prev = input[j]
				count = 1
			}
		}
		curr.WriteString(strconv.Itoa(count) + string(prev))
		input = curr.String()
	}
	return len(input)
}
