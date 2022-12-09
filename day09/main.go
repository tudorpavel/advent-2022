package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Point struct {
	X int
	Y int
}

func (p *Point) add(other Point) {
	p.X += other.X
	p.Y += other.Y
}

func (p *Point) follow(other Point) {
	diff := other.diff(*p)

	// the points are touching
	if abs(diff.X) < 2 && abs(diff.Y) < 2 {
		return
	}

	// move only 1 position towards other
	if abs(diff.X) == 2 {
		diff.X = diff.X / 2
	}

	if abs(diff.Y) == 2 {
		diff.Y = diff.Y / 2
	}

	p.add(diff)
}

func (p *Point) diff(other Point) Point {
	return Point{
		p.X - other.X,
		p.Y - other.Y,
	}
}

func solve(lines []string) (int, int) {
	deltas := map[string]Point{
		"U": {0, 1},
		"R": {1, 0},
		"D": {0, -1},
		"L": {-1, 0},
	}
	rope := [10]Point{}
	p1 := 1
	p2 := 1
	visitedP1 := map[Point]bool{}
	visitedP2 := map[Point]bool{}
	visitedP1[rope[1]] = true
	visitedP2[rope[9]] = true

	for _, line := range lines {
		var dir string
		var steps int

		fmt.Sscanf(line, "%s %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			// move head knot
			rope[0].add(deltas[dir])

			// move tail knots
			for i := 1; i < len(rope); i++ {
				rope[i].follow(rope[i-1])
			}

			// Part 1
			if !visitedP1[rope[1]] {
				p1++
				visitedP1[rope[1]] = true
			}

			// Part 2
			if !visitedP2[rope[9]] {
				p2++
				visitedP2[rope[9]] = true
			}
		}
	}

	return p1, p2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	p1, p2 := solve(lines)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
}
