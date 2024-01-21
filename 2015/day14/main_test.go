package main

import "testing"

func TestDeersMovementParsing(t *testing.T) {
	input = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	parsed := parseInput(input)

	firstExpected := DeerMovement{
		Name:           "Comet",
		Speed:          14,
		TimeInMovement: 10,
		TimeInRest:     127,
	}
	if parsed[0].Name != firstExpected.Name {
		t.Errorf("Error parsing deers: Name expected %s but found %s", firstExpected.Name, parsed[0].Name)
	}
	if parsed[0].Speed != firstExpected.Speed {
		t.Errorf("Error parsing deers: Speed expected %d but found %d", firstExpected.Speed, parsed[0].Speed)
	}
	if parsed[0].TimeInMovement != firstExpected.TimeInMovement {
		t.Errorf("Error parsing deers: TimeInMovement expected %d but found %d", firstExpected.TimeInMovement, parsed[0].TimeInMovement)
	}
	if parsed[0].TimeInRest != firstExpected.TimeInRest {
		t.Errorf("Error parsing deers: TimeInRest expected %d but found %d", firstExpected.TimeInRest, parsed[0].TimeInRest)
	}

	secondExpected := DeerMovement{
		Name:           "Dancer",
		Speed:          16,
		TimeInMovement: 11,
		TimeInRest:     162,
	}
	if parsed[1].Name != secondExpected.Name {
		t.Errorf("Error parsing deers: Name expected %s but found %s", secondExpected.Name, parsed[0].Name)
	}
	if parsed[1].Speed != secondExpected.Speed {
		t.Errorf("Error parsing deers: Speed expected %d but found %d", secondExpected.Speed, parsed[0].Speed)
	}
	if parsed[1].TimeInMovement != secondExpected.TimeInMovement {
		t.Errorf("Error parsing deers: TimeInMovement expected %d but found %d", secondExpected.TimeInMovement, parsed[0].TimeInMovement)
	}
	if parsed[1].TimeInRest != secondExpected.TimeInRest {
		t.Errorf("Error parsing deers: TimeInRest expected %d but found %d", secondExpected.TimeInRest, parsed[0].TimeInRest)
	}
}

func TestPart1(t *testing.T) {
	input := parseInput(`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`)
	time := 1000
	expected := 1120

	actual := part1(input, time)
	if expected != actual {
		t.Errorf("Incorrent answer: expected: %d, but actual: %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`)
	time := 1000
	expected := 689

	actual := part2(input, time)
	if expected != actual {
		t.Errorf("Incorrent answer: expected: %d, but actual: %d", expected, actual)
	}
}
