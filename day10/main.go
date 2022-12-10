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

func solve(lines []string) (int, int) {
	p1 := 0
	p2 := 0
	reg := 1
	cycle := 1
	interestingCycles := []int{20, 60, 100, 140, 180, 220}

	for _, line := range lines {
		if contains(interestingCycles, cycle) {
			p1 += cycle * reg
		}

		var op string
		var val int
		fmt.Sscanf(line, "%s %d", &op, &val)

		switch op {
		case "addx":
			// Check if the interesting cycle lands in
			// the middle of an addx operation
			if contains(interestingCycles, cycle+1) {
				p1 += (cycle + 1) * reg
			}

			reg += val
			cycle += 2
		case "noop":
			cycle++
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
