package main

import "testing"

func TestPart1(t *testing.T) {
	got := part1([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	})
	want := 157

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
