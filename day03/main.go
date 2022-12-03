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

func computeSeen(s string) []bool {
	seen := make([]bool, 53)
	for _, item := range s {
		seen[priority(item)] = true
	}
	return seen
}

func commonItem(rucksack string) int {
	seen := computeSeen(rucksack[:len(rucksack)/2])

	for _, item := range rucksack[len(rucksack)/2:] {
		if seen[priority(item)] {
			return priority(item)
		}
	}

	return -1
}

func badgeItem(r1 string, r2 string, r3 string) int {
	r1Seen := computeSeen(r1)
	r2Seen := computeSeen(r2)

	for _, item := range r3 {
		if r1Seen[priority(item)] && r2Seen[priority(item)] {
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

func part2(rucksacks []string) int {
	res := 0

	for i := 0; i < len(rucksacks); i += 3 {
		res += badgeItem(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
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
	fmt.Println("Part 2:", part2(lines))
}
