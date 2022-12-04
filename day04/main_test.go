package main

import "testing"

var exampleInput = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func parseInput(lines []string) []Pair {
	var pairs []Pair

	for _, line := range lines {
		pairs = append(pairs, newPair(line))
	}

	return pairs
}

func TestPart1(t *testing.T) {
	got := part1(parseInput(exampleInput))
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(parseInput(exampleInput))
	want := 4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
