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

type DeerMovement struct {
	Name           string
	Speed          int
	TimeInMovement int
	TimeInRest     int
}

func main() {
	deers := parseInput(input)
	time := 2503

	points := part1(deers, time)
	fmt.Printf("Part 1: %d\n", points)

	points = part2(deers, time)
	fmt.Printf("Part 2: %d\n", points)
}

func parseInput(input string) []DeerMovement {
	lines := strings.Split(input, "\n")
	result := make([]DeerMovement, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		speed, err := strconv.Atoi(words[3])
		if err != nil {
			panic(err)
		}

		timeInMovement, err := strconv.Atoi(words[6])
		if err != nil {
			panic(err)
		}

		timeInRest, err := strconv.Atoi(words[13])
		if err != nil {
			panic(err)
		}

		result[i] = DeerMovement{
			Name:           words[0],
			Speed:          speed,
			TimeInMovement: timeInMovement,
			TimeInRest:     timeInRest,
		}
	}

	return result
}

func part1(input []DeerMovement, time int) int {
	maxDistance := 0
	for _, deer := range input {
		distance := distance(deer, time)
		maxDistance = max(distance, maxDistance)
	}
	return maxDistance
}

func part2(input []DeerMovement, time int) int {
	points := make(map[int]int)
	for i := 1; i <= time; i++ {
		distances := make(map[int]int)
		for j, deer := range input {
			distances[j] = distance(deer, i)
		}

		indexes := make([]int, len(input))
		indexesCount := 0
		max := 0
		for index, distance := range distances {
			if distance > max {
				max = distance
				indexes[0] = index
				indexesCount = 1
			} else if distance == max {
				indexes[indexesCount] = index
				indexesCount++
			}
		}

		for _, index := range indexes[:indexesCount] {
			if _, ok := points[index]; !ok {
				points[index] = 1
			} else {
				points[index]++
			}
		}
	}

	return slices.Max(maps.Values(points))
}

func distance(deer DeerMovement, finishTime int) int {
	cycle := deer.TimeInMovement + deer.TimeInRest
	fullCycles := finishTime / cycle
	lastCycle := finishTime % cycle
	if lastCycle >= deer.TimeInMovement {
		return deer.Speed * deer.TimeInMovement * (fullCycles + 1)
	}
	return deer.Speed*deer.TimeInMovement*fullCycles + deer.Speed*lastCycle
}
