package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkTree(lines []string, row int, col int) (bool, int) {
	n := len(lines)
	m := len(lines[0])

	if row == 0 || col == 0 || row == n-1 || col == m-1 {
		return true, 0
	}

	val := lines[row][col]
	scenicScore := 1

	leftVisible := true
	rightVisible := true
	upVisible := true
	downVisible := true

	// go left
	for j := col - 1; j >= 0; j-- {
		if lines[row][j] >= val {
			leftVisible = false
			scenicScore *= col - j
			break
		}

		// count trees until the edge
		if j == 0 {
			scenicScore *= col
		}
	}

	// go right
	for j := col + 1; j < m; j++ {
		if lines[row][j] >= val {
			rightVisible = false
			scenicScore *= j - col
			break
		}

		// count trees until the edge
		if j == m-1 {
			scenicScore *= j - col
		}
	}

	// go up
	for i := row - 1; i >= 0; i-- {
		if lines[i][col] >= val {
			upVisible = false
			scenicScore *= row - i
			break
		}

		// count trees until the edge
		if i == 0 {
			scenicScore *= row
		}
	}

	// go down
	for i := row + 1; i < n; i++ {
		if lines[i][col] >= val {
			downVisible = false
			scenicScore *= i - row
			break
		}

		// count trees until the edge
		if i == n-1 {
			scenicScore *= i - row
		}
	}

	isVisible := leftVisible || rightVisible ||
		upVisible || downVisible

	return isVisible, scenicScore
}

func solve(lines []string) (int, int) {
	p1 := 0
	p2 := 0

	for i, row := range lines {
		for j := range row {
			visible, scenicScore := checkTree(lines, i, j)

			if visible {
				p1++
			}

			if scenicScore > p2 {
				p2 = scenicScore
			}
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
