package main

import (
	"strconv"
	"strings"
)

type Insctruction interface {
	Result() WireOperand
}

type BaseInsctruction struct {
	Output WireOperand
}

func (instruction BaseInsctruction) Result() WireOperand {
	return instruction.Output
}

type ValueInsctruction struct {
	BaseInsctruction
	Value Operand
}

type NotInsctruction struct {
	BaseInsctruction
	Operand Operand
}

type AndInsctruction struct {
	BaseInsctruction
	Left  Operand
	Right Operand
}

type OrInsctruction struct {
	BaseInsctruction
	Left  Operand
	Right Operand
}

type LeftShiftInsctruction struct {
	BaseInsctruction
	Left  Operand
	Right Operand
}

type RightShiftInsctruction struct {
	BaseInsctruction
	Left  Operand
	Right Operand
}

func parseInsctruction(line string) Insctruction {
	parts := strings.Split(line, " -> ")
	left := parts[0]
	switch {
	case strings.Contains(left, "NOT "):
		var operand Operand
		if v, err := strconv.Atoi(strings.TrimPrefix(left, "NOT ")); err == nil {
			operand = ValueOperand{
				Value: uint16(v),
			}
		} else {
			operand = WireOperand{
				Name: strings.TrimPrefix(left, "NOT "),
			}
		}
		return NotInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Operand: operand,
		}
	case strings.Contains(left, " AND "):
		operands := strings.Split(left, " AND ")
		leftOperand, rightOperand := parseTwoOperands(operands)

		return AndInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Left:  leftOperand,
			Right: rightOperand,
		}
	case strings.Contains(left, " OR "):
		operands := strings.Split(left, " OR ")
		leftOperand, rightOperand := parseTwoOperands(operands)

		return OrInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Left:  leftOperand,
			Right: rightOperand,
		}
	case strings.Contains(left, " LSHIFT "):
		operands := strings.Split(left, " LSHIFT ")
		leftOperand, rightOperand := parseTwoOperands(operands)

		return LeftShiftInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Left:  leftOperand,
			Right: rightOperand,
		}
	case strings.Contains(left, " RSHIFT "):
		operands := strings.Split(left, " RSHIFT ")
		leftOperand, rightOperand := parseTwoOperands(operands)

		return RightShiftInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Left:  leftOperand,
			Right: rightOperand,
		}
	default:
		var operand Operand
		if v, err := strconv.Atoi(left); err == nil {
			operand = ValueOperand{
				Value: uint16(v),
			}
		} else {
			operand = WireOperand{
				Name: left,
			}
		}
		return ValueInsctruction{
			BaseInsctruction: BaseInsctruction{
				Output: WireOperand{
					Name: parts[1],
				},
			},
			Value: operand,
		}
	}
}
