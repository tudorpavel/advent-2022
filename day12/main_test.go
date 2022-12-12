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

var inputLines = strings.Split(exampleInput, "\n")

func TestPart1(t *testing.T) {
	got, _ := solve(inputLines)
	want := 31

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	_, got := solve(inputLines)
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
