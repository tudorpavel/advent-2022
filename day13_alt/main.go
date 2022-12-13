package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"time"
)

func compare(left []any, right []any) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]

		lIsNum := reflect.TypeOf(l).Name() == "float64"
		rIsNum := reflect.TypeOf(r).Name() == "float64"

		if lIsNum && rIsNum {
			if l.(float64) < r.(float64) {
				return -1
			}
			if l.(float64) > r.(float64) {
				return 1
			}
		} else {
			var lChildren []any
			var rChildren []any

			if lIsNum {
				lChildren = []any{l}
			} else {
				lChildren = l.([]any)
			}

			if rIsNum {
				rChildren = []any{r}
			} else {
				rChildren = r.([]any)
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

func less(left []any, right []any) bool {
	return compare(left, right) < 1
}

func solve(lines []string) (int, int) {
	packets := [][]any{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var packet []any
		json.Unmarshal([]byte(line), &packet)
		packets = append(packets, packet)
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
	var div1 []any
	var div2 []any
	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)
	packets = append(packets, div1, div2)

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
		if compare(packet, div1) == 0 || compare(packet, div2) == 0 {
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
