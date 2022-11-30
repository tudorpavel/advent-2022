package main

import "testing"

func TestPart1(t *testing.T) {
	got := part1([]string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	})
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
