package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Pos struct {
	i int
	j int
}

type Cave struct {
	rocks  [][]bool
	maxRow int
}

func NewCave() Cave {
	c := Cave{}
	c.rocks = make([][]bool, 500)
	for i := range c.rocks {
		c.rocks[i] = make([]bool, 1000)
	}
	return c
}

func (cave *Cave) addRocks(points []string) {
	for k := 0; k < len(points)-1; k++ {
		var start Pos
		var end Pos
		fmt.Sscanf(points[k], "%d,%d", &start.j, &start.i)
		fmt.Sscanf(points[k+1], "%d,%d", &end.j, &end.i)

		if start.i > cave.maxRow {
			cave.maxRow = start.i
		}
		if end.i > cave.maxRow {
			cave.maxRow = end.i
		}

		var s, e int
		var isVertical bool
		if start.i == end.i {
			s = start.j
			e = end.j
			isVertical = false
		} else {
			s = start.i
			e = end.i
			isVertical = true
		}
		for s != e {
			var p Pos
			if isVertical {
				p = Pos{s, start.j}
			} else {
				p = Pos{start.i, s}
			}
			cave.rocks[p.i][p.j] = true

			if s < e {
				s++
			} else {
				s--
			}
		}
		cave.rocks[end.i][end.j] = true
	}
}

func (cave *Cave) addSand(part2 bool) bool {
	curr := Pos{0, 500}

	// All filled up with Part 2 floor
	if cave.rocks[curr.i][curr.j] {
		return false
	}

	for part2 || curr.i < cave.maxRow {
		left := Pos{curr.i + 1, curr.j - 1}
		down := Pos{curr.i + 1, curr.j}
		right := Pos{curr.i + 1, curr.j + 1}

		rockLeft := cave.rocks[left.i][left.j]
		rockDown := cave.rocks[down.i][down.j]
		rockRight := cave.rocks[right.i][right.j]

		// Part 2 infinite floor
		if down.i == cave.maxRow+2 {
			rockLeft = true
			rockDown = true
			rockRight = true
		}

		if !rockDown {
			curr.i++
			continue
		}
		if !rockLeft {
			curr.i++
			curr.j--
			continue
		}
		if !rockRight {
			curr.i++
			curr.j++
			continue
		}

		// final sand resting place
		cave.rocks[curr.i][curr.j] = true
		return true
	}

	return false
}

func solve(lines []string) (int, int) {
	cave := NewCave()

	for _, line := range lines {
		points := strings.Split(line, " -> ")
		cave.addRocks(points)
	}

	// Part 1
	p1 := 0
	for cave.addSand(false) {
		p1++
	}

	// Part 2
	p2 := p1 // continue where Part 1 left off
	for cave.addSand(true) {
		p2++
	}

	return p1, p2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := time.Now()
	p1, p2 := solve(lines)
	elapsed := time.Since(start)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
	fmt.Println("Execution time:", elapsed)
}
