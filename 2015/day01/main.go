package main

import (
	_ "embed"
	"errors"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	floor, err := part1(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Part 1: %d\n", floor)
	}

	position, err := part2(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Part 2: %d\n", position)
	}
}

func part1(input string) (floor int, _ error) {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			return floor, fmt.Errorf("Unexpected input value: %c", input[i])
		}
	}
	return floor, nil
}

func part2(input string) (int, error) {
	floor := 0
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			return -1, fmt.Errorf("Unexpected input value: %c", input[i])
		}
		if floor < 0 {
			return i + 1, nil
		}
	}
	return -1, errors.New("Santa never enters the basement")
}
