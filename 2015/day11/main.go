package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	next := part1(input)
	fmt.Printf("Part 1: %s\n", next)

	next = part1(next)
	fmt.Printf("Part 2: %s\n", next)
}

func part1(input string) string {
	for {
		var err error
		input, err = nextPassword(input, len(input)-1)
		if err != nil {
			panic(err)
		}

		if meetRequirements(input) {
			return input
		}
	}
}

func nextPassword(password string, n int) (string, error) {
	if n == -1 {
		return "", fmt.Errorf("can't find next password")
	}
	var newPassword string
	if password[n] == 'z' {
		var err error
		newPassword, err = nextPassword(password, n-1)
		if err != nil {
			return "", err
		}
		out := []byte(newPassword)
		out[n] = 'a'
		return string(out), nil
	} else {
		newPassword = password
	}

	out := []byte(newPassword)
	out[n] = newPassword[n] + 1
	return string(out), nil
}

func meetRequirements(password string) bool {
	notAllowed := "ilo"
	countOfDoubles := 0
	skipDoublesCheck := false
	hasTripleIncrease := false
	for index, letter := range password {
		for _, notAllowedLetter := range notAllowed {
			if letter == notAllowedLetter {
				return false
			}
		}

		if !skipDoublesCheck {
			if index+1 < len(password) && password[index] == password[index+1] {
				skipDoublesCheck = true
				countOfDoubles++
			}
		} else {
			skipDoublesCheck = false
		}

		if !hasTripleIncrease && index+2 < len(password) && password[index]+1 == password[index+1] && password[index+2] == password[index+1]+1 {
			hasTripleIncrease = true
		}
	}

	return countOfDoubles >= 2 && hasTripleIncrease
}
