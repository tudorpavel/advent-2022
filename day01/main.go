package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(lines []string) int {
	result := 0
	prev := -1

	for _, line := range lines {
		i, _ := strconv.Atoi(line)

		if prev != -1 && prev < i {
			result++
		}

		prev = i
	}

	return result
}

// AoC 2021 - Day 1
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", part1(lines))
}
