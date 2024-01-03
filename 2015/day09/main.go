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
	fmt.Printf("Part 2: %d\n", count)
}

func part1(input string) int {
	distances := make(map[string]map[string]int)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		locations := strings.Split(parts[0], " to ")
		if v, err := strconv.Atoi(parts[1]); err == nil {
			if _, ok := distances[locations[0]]; !ok {
				distances[locations[0]] = make(map[string]int)
			}
			if _, ok := distances[locations[1]]; !ok {
				distances[locations[1]] = make(map[string]int)
			}
			distances[locations[0]][locations[1]] = v
			distances[locations[1]][locations[0]] = v
		}
	}

	visited := make(map[string]bool)

	var fun func(string) int
	fun = func(next string) int {
		visited[next] = true
		if len(visited) == len(distances) {
			delete(visited, next)
			return 0
		}
		min := int(^uint(0) >> 1)
		for key, value := range distances {
			if _, ok := visited[key]; !ok {
				m := fun(key)
				if min > m+value[next] {
					min = m + value[next]
				}
			}
		}
		delete(visited, next)
		return min
	}

	result := int(^uint(0) >> 1)
	for name := range distances {
		result = min(result, fun(name))
	}
	return result
}

func part2(input string) int {
	distances := make(map[string]map[string]int)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		locations := strings.Split(parts[0], " to ")
		if v, err := strconv.Atoi(parts[1]); err == nil {
			if _, ok := distances[locations[0]]; !ok {
				distances[locations[0]] = make(map[string]int)
			}
			if _, ok := distances[locations[1]]; !ok {
				distances[locations[1]] = make(map[string]int)
			}
			distances[locations[0]][locations[1]] = v
			distances[locations[1]][locations[0]] = v
		}
	}

	visited := make(map[string]bool)

	var fun func(string) int
	fun = func(next string) int {
		visited[next] = true
		if len(visited) == len(distances) {
			delete(visited, next)
			return 0
		}
		max := 0
		for key, value := range distances {
			if _, ok := visited[key]; !ok {
				m := fun(key)
				if max < m+value[next] {
					max = m + value[next]
				}
			}
		}
		delete(visited, next)
		return max
	}

	result := 0
	for name := range distances {
		result = max(result, fun(name))
	}
	return result
}
