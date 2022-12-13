package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func less(left *Node, right *Node) bool {
	return compare(left.children, right.children) < 1
}

func solve(lines []string) (int, int) {
	packets := []*Node{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		packets = append(packets, NewNode(line))
	}

	// Part 1
	p1 := 0
	for i := 0; i < len(packets); i += 2 {
		if less(packets[i], packets[i+1]) {
			p1 += (i / 2) + 1
		}
	}

	// Part 2
	// Add divider packets
	div1 := NewNode("[[2]]")
	div2 := NewNode("[[6]]")
	packets = append(packets, div1)
	packets = append(packets, div2)

	// Sort packets
	sort.Slice(
		packets,
		func(i, j int) bool {
			return less(packets[i], packets[j])
		},
	)

	// Compute decoder key
	p2 := 1
	for i, packet := range packets {
		if packet == div1 || packet == div2 {
			p2 *= i + 1
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

	start := time.Now()
	p1, p2 := solve(lines)
	elapsed := time.Since(start)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
	fmt.Println("Execution time:", elapsed)
}
