package main

import (
	"reflect"
	"testing"
)

func TestSortCalories(t *testing.T) {
	got := sortCalories([]string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	})
	want := []int{
		24000,
		11000,
		10000,
		6000,
		4000,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	got := part1([]int{
		24000,
		11000,
		10000,
		6000,
		4000,
	})
	want := 24000

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2([]int{
		24000,
		11000,
		10000,
		6000,
		4000,
	})
	want := 45000

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
