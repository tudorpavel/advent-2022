package main

import (
	"bufio"
	"fmt"
	"os"
)

// Stack implementation borrowed from:
// https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(val byte) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Return top element without removing it
func (s *Stack) Peek() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func parseInput(lines []string) ([]Stack, []string) {
	// We can compute stackCount based on length and format of input lines
	// [Z] [M] [P]
	// 01234567890
	// 12 / 4 = 3
	stackCount := (len(lines[0]) + 1) / 4
	stacks := make([]Stack, stackCount)
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
				stacks[j].Push(crate)
			}
		}
	}

	return stacks, lines[separatorIndex+1:]
}

func computeResult(stacks []Stack) string {
	var res []byte

	for _, stack := range stacks {
		val, _ := stack.Peek()
		res = append(res, val)
	}

	return string(res)
}

func part1(lines []string) string {
	stacks, operations := parseInput(lines)

	// Perform rearrangement procedure
	for _, op := range operations {
		var count, from, to int
		fmt.Sscanf(op, "move %d from %d to %d", &count, &from, &to)
		fromStack := &stacks[from-1]
		toStack := &stacks[to-1]

		for i := 0; i < count; i++ {
			val, _ := fromStack.Pop()
			toStack.Push(val)
		}
	}

	return computeResult(stacks)
}

func part2(lines []string) string {
	stacks, operations := parseInput(lines)

	// Perform rearrangement procedure
	for _, op := range operations {
		var count, from, to int
		fmt.Sscanf(op, "move %d from %d to %d", &count, &from, &to)
		fromStack := &stacks[from-1]
		toStack := &stacks[to-1]

		// Use an intermediary stack to preserve the order
		// when moving multiple elements
		var temp Stack

		for i := 0; i < count; i++ {
			val, _ := fromStack.Pop()
			temp.Push(val)
		}

		for {
			val, ok := temp.Pop()

			if !ok {
				break
			}

			toStack.Push(val)
		}
	}

	return computeResult(stacks)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part1:", part1(lines))
	fmt.Println("Part2:", part2(lines))
}
