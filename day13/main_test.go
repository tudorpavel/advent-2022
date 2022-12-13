package main

import (
	"strings"
	"testing"
)

var exampleInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
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
	want := -2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
