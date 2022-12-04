package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func newRange(s string) Range {
	res := Range{}
	vals := strings.Split(s, "-")

	res.Min, _ = strconv.Atoi(vals[0])
	res.Max, _ = strconv.Atoi(vals[1])

	return res
}

type Pair struct {
	Left  Range
	Right Range
}

func newPair(s string) Pair {
	res := Pair{}
	vals := strings.Split(s, ",")

	res.Left = newRange(vals[0])
	res.Right = newRange(vals[1])

	return res
}

func areNested(r1 Range, r2 Range) bool {
	return r1.Min <= r2.Min && r2.Max <= r1.Max
}

func hasOverlap(pair Pair) bool {
	maxOfMins := -1
	minOfMaxes := -1

	if pair.Left.Min < pair.Right.Min {
		maxOfMins = pair.Right.Min
	} else {
		maxOfMins = pair.Left.Min
	}

	if pair.Left.Max < pair.Right.Max {
		minOfMaxes = pair.Left.Max
	} else {
		minOfMaxes = pair.Right.Max
	}

	return maxOfMins <= minOfMaxes
}

func part1(pairs []Pair) int {
	res := 0

	for _, pair := range pairs {
		if areNested(pair.Left, pair.Right) ||
			areNested(pair.Right, pair.Left) {
			res++
		}
	}

	return res
}

func part2(pairs []Pair) int {
	res := 0

	for _, pair := range pairs {
		if hasOverlap(pair) {
			res++
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pairs []Pair

	for scanner.Scan() {
		pairs = append(pairs, newPair(scanner.Text()))
	}

	fmt.Println("Part 1:", part1(pairs))
	fmt.Println("Part 2:", part2(pairs))
}
