package main

import (
	_ "embed"
	"fmt"

	"golang.org/x/exp/maps"
)

type Point struct {
	X int
	Y int
}

//go:embed input.txt
var input string

func main() {
	visitedBySanta := part1(input)
	fmt.Printf("Part 1: %d\n", visitedBySanta)

	visitedBySantaAndRobot := part2(input)
	fmt.Printf("Part 2: %d\n", visitedBySantaAndRobot)
}

func part2(input string) int {
	santa := Point{0, 0}
	robot := Point{0, 0}
	visitedHouses := make(map[Point]struct{})
	visitedHouses[santa] = struct{}{}
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			switch input[i] {
			case '^':
				santa.Y++
			case 'v':
				santa.Y--
			case '>':
				santa.X++
			case '<':
				santa.X--
			}
			visitedHouses[santa] = struct{}{}
		} else {
			switch input[i] {
			case '^':
				robot.Y++
			case 'v':
				robot.Y--
			case '>':
				robot.X++
			case '<':
				robot.X--
			}
			visitedHouses[robot] = struct{}{}
		}

	}

	return len(maps.Keys(visitedHouses))
}

func part1(input string) int {
	position := Point{0, 0}
	visitedHouses := make(map[Point]struct{})
	visitedHouses[position] = struct{}{}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '^':
			position.Y++
		case 'v':
			position.Y--
		case '>':
			position.X++
		case '<':
			position.X--
		}
		visitedHouses[position] = struct{}{}
	}

	return len(maps.Keys(visitedHouses))
}
