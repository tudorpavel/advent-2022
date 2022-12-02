package main

import "testing"

func TestPart1(t *testing.T) {
	got := part1([]string{
		"A Y",
		"B X",
		"C Z",
	})
	want := 15

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2([]string{
		"A Y",
		"B X",
		"C Z",
	})
	want := 12

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
