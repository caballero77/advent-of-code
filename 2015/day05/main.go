package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input      string
	vowels     = [...]byte{'a', 'e', 'i', 'o', 'u'}
	disallowed = [...]string{"ab", "cd", "pq", "xy"}
)

func main() {
	niceCount := part1(input)
	fmt.Printf("Part 1: %d\n", niceCount)

	niceCount = part2(input)
	fmt.Printf("Part 2: %d\n", niceCount)
}

func isVowel(letter byte) bool {
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == letter {
			return true
		}
	}
	return false
}

func isDisallowed(s string) bool {
	for i := 0; i < len(disallowed); i++ {
		if disallowed[i] == s {
			return true
		}
	}
	return false
}

func isNicePart1(line string) bool {
	vowelsCount := 0
	containsPair := false
	containsDisallowed := false
	for i := 0; i < len(line); i++ {
		if vowelsCount < 3 && isVowel(line[i]) {
			vowelsCount++
		}
		if i > 0 {
			if !containsPair && line[i] == line[i-1] {
				containsPair = true
			}
			if !containsDisallowed && isDisallowed(line[i-1:i+1]) {
				return false
			}
		}
	}

	return vowelsCount >= 3 && containsPair
}

func part1(input string) int {
	count := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		if isNicePart1(lines[i]) {
			count++
		}
	}
	return count
}

func isNicePart2(line string) bool {
	pairs := make(map[string]int)
	containsTwoNotOverlappingPairs := false
	containsRepeatedLetterWithOneBetween := false
	for i := 1; i < len(line); i++ {
		if pos, ok := pairs[line[i-1:i+1]]; ok {
			if i >= pos+2 {
				containsTwoNotOverlappingPairs = true
			}
		} else {
			pairs[line[i-1:i+1]] = i
		}
		if i >= 2 {
			if line[i] == line[i-2] {
				containsRepeatedLetterWithOneBetween = true
			}
		}
	}

	return containsTwoNotOverlappingPairs && containsRepeatedLetterWithOneBetween
}

func part2(input string) int {
	count := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		if isNicePart2(lines[i]) {
			count++
		}
	}
	return count
}
