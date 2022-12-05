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

func solve(lines []string) string {
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
			crate := lines[i][j*4+1]

			if crate != ' ' {
				stacks[j].Push(crate)
			}
		}
	}

	// Perform rearrangement procedure
	for _, line := range lines[separatorIndex+1:] {
		var count, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		fromStack := &stacks[from-1]
		toStack := &stacks[to-1]

		for i := 0; i < count; i++ {
			val, _ := fromStack.Pop()
			toStack.Push(val)
		}
	}

	// Compute results
	var p1 []byte

	for _, stack := range stacks {
		val, _ := stack.Peek()
		p1 = append(p1, val)
	}

	return string(p1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	p1 := solve(lines)

	fmt.Println("Part1:", p1)
}
