package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Box struct {
	Length int
	Width  int
	Heigth int
}

func (box Box) GetRequiredPaperArea() int {
	first := box.Length * box.Heigth
	second := box.Heigth * box.Width
	third := box.Width * box.Length

	return 2*first + 2*second + 2*third + min(first, second, third)
}

func (box Box) GetRequiredRibbonLength() int {
	return (box.Heigth * box.Length * box.Width) + (2*box.Heigth + 2*box.Length + 2*box.Width - 2*max(box.Heigth, box.Length, box.Width))
}

func main() {
	boxes, err := serializeInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	paperArea := part1(boxes)
	fmt.Printf("Part 1: %d\n", paperArea)

	ribbonLength := part2(boxes)
	fmt.Printf("Part 2: %d\n", ribbonLength)

}

func part1(boxes []Box) int {
	result := 0
	for i := 0; i < len(boxes); i++ {
		result += boxes[i].GetRequiredPaperArea()
	}

	return result
}

func part2(boxes []Box) int {
	result := 0
	for i := 0; i < len(boxes); i++ {
		result += boxes[i].GetRequiredRibbonLength()
	}

	return result
}

func serializeInput(input string) ([]Box, error) {
	lines := strings.Split(input, "\n")
	boxes := make([]Box, len(lines))
	for i := 0; i < len(lines); i++ {
		size := strings.Split(lines[i], "x")
		length, err := strconv.Atoi(size[0])
		if err != nil {
			return nil, err
		}

		width, err := strconv.Atoi(size[1])
		if err != nil {
			return nil, err
		}

		heigth, err := strconv.Atoi(size[2])
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, Box{length, width, heigth})
	}
	return boxes, nil
}
