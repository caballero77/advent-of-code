package main

import (
	"strconv"
)

type OperandType string

const (
	WireOperandType  OperandType = "wire"
	ValueOperandType OperandType = "value"
)

type Operand interface {
	Type() OperandType
}

type WireOperand struct {
	Name string
}

func (_ WireOperand) Type() OperandType {
	return WireOperandType
}

type ValueOperand struct {
	Value uint16
}

func (_ ValueOperand) Type() OperandType {
	return ValueOperandType
}

func parseTwoOperands(operands []string) (Operand, Operand) {
	var leftOperand Operand
	var rightOperand Operand

	if v, err := strconv.Atoi(operands[0]); err == nil {
		leftOperand = ValueOperand{
			Value: uint16(v),
		}
	} else {
		leftOperand = WireOperand{
			Name: operands[0],
		}
	}

	if v, err := strconv.Atoi(operands[1]); err == nil {
		rightOperand = ValueOperand{
			Value: uint16(v),
		}
	} else {
		rightOperand = WireOperand{
			Name: operands[1],
		}
	}

	return leftOperand, rightOperand
}
