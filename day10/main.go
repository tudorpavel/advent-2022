package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(ints []int, val int) bool {
	for _, n := range ints {
		if n == val {
			return true
		}
	}

	return false
}

func draw(cycle int, reg int, crtScreen [6][]rune) {
	col := (cycle - 1) % 40
	row := ((cycle - 1) / 40) % 6

	if reg-1 <= col && col <= reg+1 {
		crtScreen[row][col] = '#'
	} else {
		crtScreen[row][col] = '.'
	}
}

func solve(lines []string) (int, [6]string) {
	p1 := 0
	reg := 1
	cycle := 1
	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	crtScreen := [6][]rune{
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
		[]rune("........................................"),
	}

	for _, line := range lines {
		// Part 1
		if contains(interestingCycles, cycle) {
			p1 += cycle * reg
		}

		// Part 2
		draw(cycle, reg, crtScreen)

		var op string
		var val int
		fmt.Sscanf(line, "%s %d", &op, &val)

		switch op {
		case "addx":
			// Part 1
			// Check if the interesting cycle lands in
			// the middle of an addx operation
			if contains(interestingCycles, cycle+1) {
				p1 += (cycle + 1) * reg
			}

			// Part 2
			draw(cycle+1, reg, crtScreen)

			reg += val
			cycle += 2
		case "noop":
			cycle++
		}
	}

	p2 := [6]string{}

	for i, row := range crtScreen {
		p2[i] = string(row)
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
	fmt.Println("Part2:")
	for _, row := range p2 {
		fmt.Println(row)
	}
}
