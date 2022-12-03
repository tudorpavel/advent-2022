package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(letter rune) int {
	if letter >= 'a' {
		return int(letter - 'a' + 1)
	}
	return int(letter - 'A' + 27)
}

func commonItem(rucksack string) int {
	seen := make([]bool, 53)

	for _, item := range rucksack[:len(rucksack)/2] {
		seen[priority(item)] = true
	}

	for _, item := range rucksack[len(rucksack)/2:] {
		if seen[priority(item)] {
			return priority(item)
		}
	}

	return -1
}

func part1(rucksacks []string) int {
	res := 0

	for _, rucksack := range rucksacks {
		res += commonItem(rucksack)
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", part1(lines))
}
