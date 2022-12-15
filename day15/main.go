package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type Point struct {
	x int
	y int
}

func (p Point) distance(o Point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func (p Point) eql(o Point) bool {
	return p.x == o.x && p.y == o.y
}

type Pair struct {
	sensor Point
	beacon Point
}

func BuildPair(line string) Pair {
	res := Pair{}

	fmt.Sscanf(
		line,
		"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&res.sensor.x, &res.sensor.y,
		&res.beacon.x, &res.beacon.y,
	)

	return res
}

func solve(lines []string, checkY int) (int, int) {
	pairs := []Pair{}

	for _, line := range lines {
		pairs = append(pairs, BuildPair(line))
	}

	p1 := 0

	for x := -10000000; x < 10000000; x++ {
		candidate := Point{x, checkY}
		for _, pair := range pairs {
			minDistance := pair.sensor.distance(pair.beacon)
			if candidate.distance(pair.sensor) <= minDistance &&
				!candidate.eql(pair.beacon) {
				p1++
				break
			}
		}
	}

	return p1, -2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := time.Now()
	p1, p2 := solve(lines, 2000000)
	elapsed := time.Since(start)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
	fmt.Println("Execution time:", elapsed)
}
