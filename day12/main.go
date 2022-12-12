package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Pos struct {
	i    int
	j    int
	step int
}

func (p *Pos) eql(o Pos) bool {
	return p.i == o.i && p.j == o.j
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

func solve(lines []string) (int, int) {
	visited := make([][]bool, len(lines))
	start := Pos{}
	end := Pos{}

	for i, line := range lines {
		visited[i] = make([]bool, len(line))

		for j, r := range line {
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

	p1 := 0
	p2 := len(lines) * len(lines[0])

	// BFS
	queue := []Pos{end}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if visited[curr.i][curr.j] {
			continue
		}

		visited[curr.i][curr.j] = true

		if lines[curr.i][curr.j] == 'a' {
			if curr.eql(start) {
				p1 = curr.step
			}

			if curr.step < p2 {
				p2 = curr.step
			}

			continue
		}

		for _, delta := range [4]Pos{{i: -1}, {i: 1}, {j: -1}, {j: 1}} {
			next := delta.add(curr)
			next.step = curr.step + 1

			if next.outOfBounds(len(lines), len(lines[0])) {
				continue
			}

			c := lines[curr.i][curr.j]
			n := lines[next.i][next.j]
			// Next elevation is lower by more than 1
			if c > n && c-n > 1 {
				continue
			}

			queue = append(queue, next)
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

	start := time.Now()
	p1, p2 := solve(lines)
	elapsed := time.Since(start)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
	fmt.Println("Execution time:", elapsed)
}
