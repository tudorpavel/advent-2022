package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	val      int
	parent   *Node
	children []*Node
}

func (n *Node) isList() bool {
	return n.val == -1
}

func NewNode(line string) *Node {
	curr := &Node{val: -1}

	for len(line) > 0 {
		switch {
		case line[0] == '[':
			new := &Node{val: -1, parent: curr}
			curr.children = append(curr.children, new)
			curr = new
			line = line[1:]
		case line[0] == ']':
			curr = curr.parent
			line = line[1:]
			if len(line) > 0 && line[0] == ',' {
				line = line[1:]
			}
		default: // number
			f := func(c rune) bool {
				return c == ']' || c == ','
			}
			idx := strings.IndexFunc(line, f)
			num, _ := strconv.Atoi(line[:idx])
			curr.children = append(curr.children, &Node{val: num})
			line = line[idx:]
			if line[0] == ',' {
				line = line[1:]
			}
		}
	}

	res := curr.children[0]
	res.parent = nil

	return res
}

func compare(left []*Node, right []*Node) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]

		if !l.isList() && !r.isList() {
			if l.val < r.val {
				return -1
			}
			if l.val > r.val {
				return 1
			}
		} else {
			lChildren := l.children
			rChildren := r.children

			if !l.isList() {
				lChildren = []*Node{l}
			}

			if !r.isList() {
				rChildren = []*Node{r}
			}

			res := compare(lChildren, rChildren)

			if res != 0 {
				return res
			}
		}
	}

	if len(left) < len(right) {
		return -1
	}

	if len(left) > len(right) {
		return 1
	}

	return 0
}

func solve(lines []string) (int, int) {
	p1 := 0

	for i := 0; i < len(lines); i += 3 {
		left := NewNode(lines[i])
		right := NewNode(lines[i+1])

		if compare(left.children, right.children) < 1 {
			p1 += (i / 3) + 1
		}
	}

	return p1, -2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := time.Now()
	p1, p2 := solve(lines)
	elapsed := time.Since(start)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
	fmt.Println("Execution time:", elapsed)
}
