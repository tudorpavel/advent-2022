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

// Exclude points that are within other pairs' radius,
// it's the right beacon if it's at radius+1 from 3 other sensors
func isTheBeacon(candidate Point, pairs []Pair) bool {
	touchingEdges := make([]bool, len(pairs))

	for i, pair := range pairs {
		radius := pair.sensor.distance(pair.beacon)

		if candidate.distance(pair.sensor) <= radius {
			return false
		}

		if candidate.distance(pair.sensor) == radius+1 {
			touchingEdges[i] = true
		}
	}

	count := 0

	for _, touchingOtherEdge := range touchingEdges {
		if touchingOtherEdge {
			count++
		}
	}

	return count > 2
}

// For each sensor+beacon pair go around its diamond-shaped edge
// centered at sensor and radius distance between sensor and beacon
// and check if there's a point just outside 3 other pairs' radius
func part2(pairs []Pair) int {
	for i, pair := range pairs {
		others := []Pair{}
		others = append(others, pairs[:i]...)
		others = append(others, pairs[i+1:]...)
		dist := pair.sensor.distance(pair.beacon)

		// top left of diamond shape
		x := pair.sensor.x - dist - 1
		y := pair.sensor.y
		for x <= pair.sensor.x {
			if isTheBeacon(Point{x, y}, others) {
				return x*4000000 + y
			}

			x++
			y--
		}
		// top right of diamond shape
		x = pair.sensor.x
		y = pair.sensor.y - dist - 1
		for y <= pair.sensor.y {
			if isTheBeacon(Point{x, y}, others) {
				return x*4000000 + y
			}

			x++
			y++
		}
		// bottom right of diamond shape
		x = pair.sensor.x + dist + 1
		y = pair.sensor.y
		for x >= pair.sensor.x {
			if isTheBeacon(Point{x, y}, others) {
				return x*4000000 + y
			}

			x--
			y++
		}
		// bottom left of diamond shape
		x = pair.sensor.x
		y = pair.sensor.y + dist + 1
		for y >= pair.sensor.y {
			if isTheBeacon(Point{x, y}, others) {
				return x*4000000 + y
			}

			x--
			y--
		}
	}

	return -1
}

func solve(lines []string, checkY int) (int, int) {
	pairs := []Pair{}

	for _, line := range lines {
		pairs = append(pairs, BuildPair(line))
	}

	// Part 1
	p1 := 0
	for x := -1000000; x < 5000000; x++ {
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

	// Part 2
	p2 := part2(pairs)

	return p1, p2
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
