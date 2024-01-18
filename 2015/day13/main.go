package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

//go:embed input.txt
var input string

func main() {
	data := parseInput(input)

	happyness := part1(data)
	fmt.Printf("Part 1: %d\n", happyness)

	myHappiness := make(map[string]int)
	for person, happiness := range data {
		happiness["Me"] = 0
		myHappiness[person] = 0
	}
	data["Me"] = myHappiness

	happyness = part1(data)
	fmt.Printf("Part 2: %d\n", happyness)
}

func parseInput(input string) map[string]map[string]int {
	data := make(map[string]map[string]int)

	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")

		firstPerson := words[0]
		secondPerson := words[len(words)-1][:len(words[len(words)-1])-1]
		hapiness, err := strconv.Atoi(words[3])
		if err != nil {
			panic(err)
		}
		if words[2] == "lose" {
			hapiness *= -1
		}

		value, ok := data[firstPerson]
		if !ok {
			value = make(map[string]int)
			data[firstPerson] = value
		}

		value[secondPerson] = hapiness
	}

	return data
}

func part1(input map[string]map[string]int) int {
	if len(input) == 0 {
		return 0
	}

	var inner func([]string, int) int
	inner = func(table []string, n int) int {
		if n == len(table) {
			return calculateHappiness(table, input)
		}
		maxHappiness := 0
		for firstPerson := range input {
			if slices.Contains(table, firstPerson) {
				continue
			}
			table[n] = firstPerson
			currentMaxHappiness := inner(table, n+1)
			maxHappiness = max(maxHappiness, currentMaxHappiness)
			table[n] = ""
		}
		return maxHappiness
	}

	keys := maps.Keys(input)
	person := keys[0]
	n := len(input)
	table := make([]string, n)
	table[0] = person
	result := inner(table, 1)
	return result
}

func calculateHappiness(table []string, input map[string]map[string]int) int {
	hapiness := 0
	for i, person := range table {
		if i > 0 {
			hapiness += input[person][table[i-1]] + input[table[i-1]][person]
		}
	}
	hapiness += input[table[0]][table[len(table)-1]] + input[table[len(table)-1]][table[0]]
	return hapiness
}
