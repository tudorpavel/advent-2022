package main

import (
	"strings"
	"testing"
)

var exampleInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
var inputLines = strings.Split(exampleInput, "\n")

func TestPart1(t *testing.T) {
	got, _ := solve(inputLines)
	want := 24

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	_, got := solve(inputLines)
	want := 93

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
