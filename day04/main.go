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

func part1(lines []string) int {
	res := 0

	for _, line := range lines {
		pair := newPair(line)

		if areNested(pair.Left, pair.Right) || areNested(pair.Right, pair.Left) {
			res++
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", part1(lines))
}
