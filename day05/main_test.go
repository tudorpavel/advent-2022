package main

import "testing"

var exampleInput = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	got := part1(exampleInput)
	want := "CMZ"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(exampleInput)
	want := "MCD"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
