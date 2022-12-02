package main

import (
	"bufio"
	"fmt"
	"os"
)

func totalScore(lines []string, scoreCard [][]int) int {
	score := 0

	for _, line := range lines {
		score += scoreCard[line[0]-'A'][line[2]-'X']
	}

	return score
}

func part1(lines []string) int {
	scoreCard := [][]int{
		//X Y  Z
		{4, 8, 3}, // A
		{1, 5, 9}, // B
		{7, 2, 6}, // C
	}

	return totalScore(lines, scoreCard)
}

func part2(lines []string) int {
	scoreCard := [][]int{
		//X Y  Z
		{3, 4, 8}, // A
		{1, 5, 9}, // B
		{2, 6, 7}, // C
	}

	return totalScore(lines, scoreCard)
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
