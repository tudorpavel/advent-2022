package main

import "testing"

func TestSolve(t *testing.T) {
	got := solve([]string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	})
	want := "CMZ"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
