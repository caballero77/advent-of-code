package main

import "testing"

type requirementsTest struct {
	Password string
	Expected bool
}

func TestPart1Requirements(t *testing.T) {
	tests := []requirementsTest{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
	}

	for _, test := range tests {
		if actual := meetRequirements(test.Password); actual != test.Expected {
			t.Errorf("Output for password: \"%s\" is %v when expected to be %v", test.Password, actual, test.Expected)
		}
	}
}

type nextPasswordTest struct {
	Password string
	Next     string
}

func TestPart1NextPassword(t *testing.T) {
	tests := []nextPasswordTest{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, test := range tests {
		if actual := part1(test.Password); actual != test.Next {
			t.Errorf("Output for password: \"%s\" is %s when expected to be %s", test.Password, actual, test.Next)
		}
	}
}
