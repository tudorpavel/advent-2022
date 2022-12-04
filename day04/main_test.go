package main

import "testing"

func TestPart1(t *testing.T) {
	got := part1([]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	})
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
