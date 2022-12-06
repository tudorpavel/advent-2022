package main

import "testing"

var exampleInput = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

func TestPart1(t *testing.T) {
	got := markerIndex(exampleInput, 4)
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := markerIndex(exampleInput, 14)
	want := 29

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLimit(t *testing.T) {
	got := markerIndex("aaabcd", 4)
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
