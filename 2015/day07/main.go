package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	instructions := parseInstructions(input)
	memory := make(map[string]uint16)

	count, err := part1(memory, instructions, "a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part 1: %d\n", count)

	memory = make(map[string]uint16)

	instructions["b"] = ValueInsctruction{
		BaseInsctruction: BaseInsctruction{Output: WireOperand{Name: "b"}},
		Value:            ValueOperand{Value: count},
	}

	count, err = part1(memory, instructions, "a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part 2: %d\n", count)
}

func parseInstructions(input string) map[string]Insctruction {
	lines := strings.Split(input, "\n")
	instructions := make(map[string]Insctruction, len(lines))
	for i := 0; i < len(lines); i++ {
		instruction := parseInsctruction(lines[i])
		instructions[instruction.Result().Name] = instruction
	}
	return instructions
}

func part1(memory map[string]uint16, instructions map[string]Insctruction, result string) (uint16, error) {
	if v, ok := memory[result]; ok {
		return v, nil
	}
	instruction := instructions[result]
	var value uint16
	switch instruction := instruction.(type) {
	case ValueInsctruction:
		var err error
		value, err = getValue(memory, instructions, instruction.Value)
		if err != nil {
			return 0, err
		}
	case NotInsctruction:
		operand, err := getValue(memory, instructions, instruction.Operand)
		if err != nil {
			return 0, err
		}

		value = ^operand
	case AndInsctruction:
		left, err := getValue(memory, instructions, instruction.Left)
		if err != nil {
			return 0, err
		}

		right, err := getValue(memory, instructions, instruction.Right)
		if err != nil {
			return 0, err
		}

		value = left & right
	case OrInsctruction:
		left, err := getValue(memory, instructions, instruction.Left)
		if err != nil {
			return 0, err
		}

		right, err := getValue(memory, instructions, instruction.Right)
		if err != nil {
			return 0, err
		}

		value = left | right
	case LeftShiftInsctruction:
		left, err := getValue(memory, instructions, instruction.Left)
		if err != nil {
			return 0, err
		}

		right, err := getValue(memory, instructions, instruction.Right)
		if err != nil {
			return 0, err
		}

		value = left << right
	case RightShiftInsctruction:
		left, err := getValue(memory, instructions, instruction.Left)
		if err != nil {
			return 0, err
		}

		right, err := getValue(memory, instructions, instruction.Right)
		if err != nil {
			return 0, err
		}

		value = left >> right
	}
	memory[result] = value
	return value, nil
}

func getValue(memory map[string]uint16, instructions map[string]Insctruction, value Operand) (uint16, error) {
	switch value := value.(type) {
	case ValueOperand:
		return value.Value, nil
	case WireOperand:
		return part1(memory, instructions, value.Name)
	}
	return 0, errors.New("unknown value")
}
