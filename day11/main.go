package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items        []int
	operation    byte
	operand      int
	divisor      int
	trueMonkey   *Monkey
	falseMonkey  *Monkey
	inspectCount int
}

func (m *Monkey) pushItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) adjustWorryLevel(item int) int {
	operand := m.operand

	// Operand was "old"
	if operand == -1 {
		operand = item
	}

	if m.operation == '+' {
		item += operand
	} else {
		item *= operand
	}

	// Adjust for relief rule
	return item / 3
}

func (m *Monkey) takeTurn() {
	m.inspectCount += len(m.items)

	for len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]

		item = m.adjustWorryLevel(item)

		if item%m.divisor == 0 {
			m.trueMonkey.pushItem(item)
		} else {
			m.falseMonkey.pushItem(item)
		}
	}
}

func solve(lines []string) (int, int) {
	p2 := 0
	monkeys := []*Monkey{}

	// Init monkeys
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, &Monkey{})
	}

	// Parse input
	for i := 0; i < len(lines); i += 7 {
		monkey := monkeys[i/7]

		// Items
		for _, s := range strings.Split(lines[i+1][18:], ", ") {
			item, _ := strconv.Atoi(s)
			monkey.pushItem(item)
		}

		// Operation
		op := lines[i+2][23:]
		num, err := strconv.Atoi(op[2:])
		if err != nil {
			num = -1 // second operand is "old"
		}
		monkey.operation = op[0]
		monkey.operand = num

		// Test
		fmt.Sscanf(
			lines[i+3],
			"  Test: divisible by %d",
			&monkey.divisor)
		var j int
		fmt.Sscanf(lines[i+4], "    If true: throw to monkey %d", &j)
		monkey.trueMonkey = monkeys[j]
		fmt.Sscanf(lines[i+5], "    If false: throw to monkey %d", &j)
		monkey.falseMonkey = monkeys[j]
	}

	// Run 20 rounds for Part 1
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.takeTurn()
		}
	}

	max1 := 0
	max2 := 0

	for _, monkey := range monkeys {
		c := monkey.inspectCount

		if c > max1 {
			max2 = max1
			max1 = c
		} else if c > max2 {
			max2 = c
		}
	}

	return max1 * max2, p2
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
