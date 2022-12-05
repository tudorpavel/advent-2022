package main

import (
	"bufio"
	"fmt"
	"os"
)

// Stack implementation borrowed from:
// https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack []byte

func (s *Stack) Push(vals ...byte) {
	*s = append(*s, vals...)
}

func (s *Stack) Pop(count int) ([]byte, bool) {
	if len(*s) < count {
		return nil, false
	} else {
		index := len(*s) - count
		elements := (*s)[index:]
		*s = (*s)[:index]
		return elements, true
	}
}

func (s *Stack) Peek() (byte, bool) {
	if len(*s) < 1 {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}

func computeResult(stacks []Stack) string {
	var res []byte

	for _, stack := range stacks {
		val, _ := stack.Peek()
		res = append(res, val)
	}

	return string(res)
}

func solve(lines []string) (string, string) {
	// We can compute stackCount based on length and format of input lines
	// [Z] [M] [P]
	// 01234567890
	// 12 / 4 = 3
	stackCount := (len(lines[0]) + 1) / 4
	stacks1 := make([]Stack, stackCount)
	stacks2 := make([]Stack, stackCount)
	separatorIndex := 0

	// Find the index of the empty line that separates
	// the stacks part and the rearrangement procedure
	// part of the input
	for i, line := range lines {
		if line == "" {
			separatorIndex = i
			break
		}
	}

	// Build stacks
	// Move through lines from the bottom layer of crates up,
	// if there is a create at a stack's position, push it onto its
	// corresponding stack.
	for i := separatorIndex - 2; i >= 0; i-- {
		for j := 0; j < stackCount; j++ {
			// Crate letter index can be computed based on
			// the standard format of the input.
			// [Z] [M] [P]
			// 01234567890
			crate := lines[i][j*4+1]

			if crate != ' ' {
				stacks1[j].Push(crate)
				stacks2[j].Push(crate)
			}
		}
	}

	// Perform rearrangement procedure
	for _, op := range lines[separatorIndex+1:] {
		var count, from, to int
		fmt.Sscanf(op, "move %d from %d to %d", &count, &from, &to)

		// Part 1
		for i := 0; i < count; i++ {
			vals, _ := stacks1[from-1].Pop(1)
			stacks1[to-1].Push(vals...)
		}

		// Part 2
		vals, _ := stacks2[from-1].Pop(count)
		stacks2[to-1].Push(vals...)
	}

	return computeResult(stacks1), computeResult(stacks2)
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
