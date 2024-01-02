package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	startsWithFive := part1(input)
	fmt.Printf("Part 1: %d\n", startsWithFive)

	startsWithSix := part2(input)
	fmt.Printf("Part 1: %d\n", startsWithSix)
}

func part1(input string) int {
	number := 0
	var hash [16]byte = [16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for !strings.HasPrefix(hex.EncodeToString(hash[:]), "00000") {
		number++
		hash = md5.Sum([]byte(input + strconv.Itoa(number)))
	}

	return number
}

func part2(input string) int {
	number := 0
	var hash [16]byte = [16]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for !strings.HasPrefix(hex.EncodeToString(hash[:]), "000000") {
		number++
		hash = md5.Sum([]byte(input + strconv.Itoa(number)))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), "00000") {
			fmt.Println(hex.EncodeToString(hash[:]))
		}
	}

	return number
}
