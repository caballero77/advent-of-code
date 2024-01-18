package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	sum1 := part1(input)
	fmt.Printf("Part 1: %d\n", sum1)

	sum2 := part2(input)
	fmt.Printf("Part 2: %d\n", sum2)
}

func part1(input string) int64 {
	reg, err := regexp.Compile("-?[0-9]+")
	if err != nil {
		panic(err)
	}

	var sum int64
	for _, match := range reg.FindAllString(input, -1) {
		if number, err := strconv.Atoi(match); err == nil {
			sum += int64(number)
		}
	}
	return sum
}

func part2(input string) int64 {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)

	if err != nil {
		panic(err)
	}

	return part2Sum(data)
}

func part2Sum(data interface{}) int64 {
	var sum int64
	switch data := data.(type) {
	case map[string]interface{}:
		for _, item := range data {
			if item == "red" {
				return 0
			}
			sum += part2Sum(item)
		}
	case []interface{}:
		for _, item := range data {
			sum += part2Sum(item)
		}
	case string:
		if number, err := strconv.Atoi(data); err == nil {
			sum += int64(number)
		}
	case json.Number:
		if number, err := data.Int64(); err == nil {
			sum += number
		}
	case float64:
		sum += int64(data)
	}
	return sum
}
