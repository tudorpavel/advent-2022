package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sortCalories(lines []string) []int {
	var calories []int
	current := 0

	for _, line := range lines {
		if line == "" {
			calories = append(calories, current)

			current = 0
			continue
		}

		i, _ := strconv.Atoi(line)

		current += i
	}

	calories = append(calories, current)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	return calories
}

func part1(calories []int) int {
	return calories[0]
}

func part2(calories []int) int {
	result := 0

	for _, calorie := range calories[:3] {
		result += calorie
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	calories := sortCalories(lines)
	fmt.Println("Part 1:", part1(calories))
	fmt.Println("Part 2:", part2(calories))
}
