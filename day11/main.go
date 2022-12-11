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
	trueIndex    int
	falseIndex   int
	inspectCount int
}

func (m *Monkey) pushItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) adjustItem(item int) int {
	operand := m.operand

	// Operand was "old"
	if m.operand == -1 {
		operand = item
	}

	if m.operation == '+' {
		item += operand
	} else {
		item *= operand
	}

	return item
}

type Solution struct {
	monkeys []*Monkey
	lcm     int
}

func (sol *Solution) clone() *Solution {
	newSol := &Solution{
		monkeys: []*Monkey{},
		lcm:     sol.lcm,
	}

	for _, m := range sol.monkeys {
		newM := *m
		newSol.monkeys = append(newSol.monkeys, &newM)
	}

	return newSol
}

func (sol *Solution) part1() int {
	for i := 0; i < 20; i++ {
		sol.playOneRound(func(item int) int {
			// Adjust for relief rule from Part 1
			return item / 3
		})
	}

	return sol.monkeyBusiness()
}

func (sol *Solution) part2() int {
	for i := 0; i < 10000; i++ {
		sol.playOneRound(func(item int) int {
			// Given that all test divisors are prime numbers we can
			// multiply them to get the Least Common Multiple (LCM)
			// and use that to keep items small for Part 2
			return item % sol.lcm
		})
	}

	return sol.monkeyBusiness()
}

func (sol *Solution) playOneRound(reduceWorry func(int) int) {
	for _, m := range sol.monkeys {
		m.inspectCount += len(m.items)

		for len(m.items) > 0 {
			item := m.items[0]
			m.items = m.items[1:]

			item = m.adjustItem(item)
			item = reduceWorry(item)

			if item%m.divisor == 0 {
				sol.monkeys[m.trueIndex].pushItem(item)
			} else {
				sol.monkeys[m.falseIndex].pushItem(item)
			}
		}
	}
}

func (sol *Solution) monkeyBusiness() int {
	max1 := 0
	max2 := 0

	for _, monkey := range sol.monkeys {
		c := monkey.inspectCount

		if c > max1 {
			max2 = max1
			max1 = c
		} else if c > max2 {
			max2 = c
		}
	}

	return max1 * max2
}

func solve(lines []string) (int, int) {
	sol := Solution{lcm: 1}

	// Parse input
	for i := 0; i < len(lines); i += 7 {
		monkey := &Monkey{}
		sol.monkeys = append(sol.monkeys, monkey)

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

		// Read test divisor
		fmt.Sscanf(
			lines[i+3],
			"  Test: divisible by %d",
			&monkey.divisor)

		// Compute Lowest Common Multiple (LCM)
		sol.lcm *= monkey.divisor

		// Read monkey indexes
		var j int
		fmt.Sscanf(lines[i+4], "    If true: throw to monkey %d", &j)
		monkey.trueIndex = j
		fmt.Sscanf(lines[i+5], "    If false: throw to monkey %d", &j)
		monkey.falseIndex = j
	}

	// Clone for Part 2
	solP2 := sol.clone()

	return sol.part1(), solP2.part2()
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
