package main

import "testing"

type testCase struct {
	Json        string
	ExpectedSum int64
}

func TestPart1(t *testing.T) {
	tests := []testCase{
		{"[1,2,3]", 6},
		{`{"a":2,"b":4}`, 6},
		{"[[[3]]]", 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`{}`, 0},
		{`[]`, 0},
	}

	for _, test := range tests {
		if actual := part1(test.Json); actual != test.ExpectedSum {
			t.Errorf("Output for json: \"%s\" is %d when expected to be %d", test.Json, actual, test.ExpectedSum)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []testCase{
		{"[1,2,3]", 6},
		{`{"a":2,"b":4}`, 6},
		{"[[[3]]]", 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`{}`, 0},
		{`[]`, 0},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}

	for _, test := range tests {
		if actual := part2(test.Json); actual != test.ExpectedSum {
			t.Errorf("Output for json: \"%s\" is %d when expected to be %d", test.Json, actual, test.ExpectedSum)
		}
	}
}
