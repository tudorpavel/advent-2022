package main

import (
	"strings"
	"testing"
)

var exampleInput = `30373
25512
65332
33549
35390`
var inputLines = strings.Split(exampleInput, "\n")

func TestPart1(t *testing.T) {
	got, _ := solve(inputLines)
	want := 21

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	_, got := solve(inputLines)
	want := 8

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
