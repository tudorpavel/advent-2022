package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	i int
	j int
}

func (p *Pos) add(o Pos) Pos {
	return Pos{
		i: p.i + o.i,
		j: p.j + o.j,
	}
}

func (p *Pos) outOfBounds(n int, m int) bool {
	return p.i < 0 || p.j < 0 || p.i >= n || p.j >= m
}

func fill(lines []string, dist *[][]int, curr Pos, step int) {
	val := (*dist)[curr.i][curr.j]

	// There's another shorter path
	if val > -1 && val <= step {
		return
	}

	(*dist)[curr.i][curr.j] = step
	step++

	for _, delta := range [4]Pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		next := curr.add(delta)

		if next.outOfBounds(len(lines), len(lines[0])) {
			continue
		}

		c := lines[curr.i][curr.j]
		n := lines[next.i][next.j]
		// Next elevation is higher by more than 1
		if c < n && n-c > 1 {
			continue
		}

		fill(lines, dist, curr.add(delta), step)
	}
}

func part2(lines []string, dist *[][]int, end Pos) int {
	min := len(lines) * len(lines[0])

	for i, line := range lines {
		for j, r := range line {
			if r == 'a' {
				// A potential starting point
				start := Pos{i, j}

				// Reset distance matrix
				for i := range *dist {
					for j := range (*dist)[0] {
						(*dist)[i][j] = -1
					}
				}

				fill(lines, dist, start, 0)

				val := (*dist)[end.i][end.j]
				if val > -1 && val < min {
					min = val
				}
			}
		}
	}

	return min
}

func solve(lines []string) (int, int) {
	dist := make([][]int, len(lines))
	start := Pos{}
	end := Pos{}

	for i, line := range lines {
		dist[i] = make([]int, len(line))

		for j, r := range line {
			dist[i][j] = -1

			if r == 'S' {
				start.i = i
				start.j = j
				l := []rune(lines[i])
				l[j] = 'a'
				lines[i] = string(l)
			}
			if r == 'E' {
				end.i = i
				end.j = j
				l := []rune(lines[i])
				l[j] = 'z'
				lines[i] = string(l)
			}
		}
	}

	// Part 1
	fill(lines, &dist, start, 0)
	p1 := dist[end.i][end.j]

	// Part 2
	p2 := part2(lines, &dist, end)

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
