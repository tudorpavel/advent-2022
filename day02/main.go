package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(lines []string) int {
	score := 0

	for _, line := range lines {
		if line[2] == 'X' {
			score += 1

			if line[0] == 'A' {
				score += 3
			} else if line[0] == 'B' {
				score += 0
			} else {
				score += 6
			}
		} else if line[2] == 'Y' {
			score += 2

			if line[0] == 'A' {
				score += 6
			} else if line[0] == 'B' {
				score += 3
			} else {
				score += 0
			}
		} else {
			score += 3

			if line[0] == 'A' {
				score += 0
			} else if line[0] == 'B' {
				score += 6
			} else {
				score += 3
			}
		}
	}

	return score
}

func part2(lines []string) int {
	score := 0

	for _, line := range lines {
		if line[2] == 'X' {
			score += 0

			if line[0] == 'A' {
				score += 3
			} else if line[0] == 'B' {
				score += 1
			} else {
				score += 2
			}
		} else if line[2] == 'Y' {
			score += 3

			if line[0] == 'A' {
				score += 1
			} else if line[0] == 'B' {
				score += 2
			} else {
				score += 3
			}
		} else {
			score += 6

			if line[0] == 'A' {
				score += 2
			} else if line[0] == 'B' {
				score += 3
			} else {
				score += 1
			}
		}
	}

	return score
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
