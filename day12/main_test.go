package main

import (
	"strings"
	"testing"
)

var exampleInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	inputLines := strings.Split(exampleInput, "\n")
	got, _ := solve(inputLines)
	want := 31

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	inputLines := strings.Split(exampleInput, "\n")
	_, got := solve(inputLines)
	want := 29

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
