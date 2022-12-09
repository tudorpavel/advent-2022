package main

import (
	"strings"
	"testing"
)

var exampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
var inputLines = strings.Split(exampleInput, "\n")

func TestPart1(t *testing.T) {
	got, _ := solve(inputLines)
	want := 13

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	_, got := solve(inputLines)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2Large(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	_, got := solve(strings.Split(input, "\n"))
	want := 36

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
