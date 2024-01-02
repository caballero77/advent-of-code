package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Action string

var (
	Toggle  Action = "toggle"
	TurnOn  Action = "turn on"
	TurnOff Action = "turn off"
)

type Insctruction struct {
	Action Action
	FromX  int
	ToX    int
	FromY  int
	ToY    int
}

func main() {
	insctructions, err := parseInsctructions(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := part1(insctructions)
	fmt.Printf("Part 1: %d\n", count)

	count = part2(insctructions)
	fmt.Printf("Part 2: %d\n", count)
}

func parseInsctructions(input string) ([]Insctruction, error) {
	lines := strings.Split(input, "\n")
	instructions := make([]Insctruction, len(lines))
	for i := 0; i < len(lines); i++ {
		var err error
		instructions[i], err = parseInsctruction(lines[i])
		if err != nil {
			return nil, err
		}
	}
	return instructions, nil
}

func parseInsctruction(instruction string) (Insctruction, error) {
	var action Action
	var from, to []string
	parts := strings.Split(instruction, " ")
	switch parts[0] {
	case string(Toggle):
		action = Toggle

		from = strings.Split(parts[1], ",")
		to = strings.Split(parts[3], ",")
	case "turn":
		switch parts[1] {
		case "on":
			action = TurnOn
		case "off":
			action = TurnOff
		}
		from = strings.Split(parts[2], ",")
		to = strings.Split(parts[4], ",")
	}

	fromX, err := strconv.Atoi(from[0])
	if err != nil {
		return Insctruction{}, err
	}

	fromY, err := strconv.Atoi(from[1])
	if err != nil {
		return Insctruction{}, err
	}

	toX, err := strconv.Atoi(to[0])
	if err != nil {
		return Insctruction{}, err
	}

	toY, err := strconv.Atoi(to[1])
	if err != nil {
		return Insctruction{}, err
	}

	return Insctruction{
		Action: action,
		FromX:  fromX,
		FromY:  fromY,
		ToX:    toX,
		ToY:    toY,
	}, nil
}

func part1(insctructions []Insctruction) int {
	grid := [1000][1000]bool{}
	for _, insctruction := range insctructions {
		var action func(bool) bool
		switch insctruction.Action {
		case Toggle:
			action = func(value bool) bool { return !value }
		case TurnOff:
			action = func(_ bool) bool { return false }
		case TurnOn:
			action = func(_ bool) bool { return true }
		}
		for x := insctruction.FromX; x <= insctruction.ToX; x++ {
			for y := insctruction.FromY; y <= insctruction.ToY; y++ {
				grid[x][y] = action(grid[x][y])
			}
		}
	}
	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				count++
			}
		}
	}

	return count
}

func part2(insctructions []Insctruction) int {
	grid := [1000][1000]int{}
	for _, insctruction := range insctructions {
		var action func(int) int
		switch insctruction.Action {
		case Toggle:
			action = func(value int) int { return value + 2 }
		case TurnOff:
			action = func(value int) int {
				if value == 0 {
					return 0
				}
				return value - 1
			}
		case TurnOn:
			action = func(value int) int { return value + 1 }
		}
		for x := insctruction.FromX; x <= insctruction.ToX; x++ {
			for y := insctruction.FromY; y <= insctruction.ToY; y++ {
				grid[x][y] = action(grid[x][y])
			}
		}
	}
	var sum int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += grid[x][y]
		}
	}

	return sum
}
