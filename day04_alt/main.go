package main

import (
	"fmt"
)

// Shorter implementation inspired by a C solution on Reddit:
// https://old.reddit.com/r/adventofcode/comments/zc0zta/2022_day_4_solutions/iyui6f4/
func main() {
	var p1, p2, a, b, x, y int

	for {
		_, err := fmt.Scanf("%d-%d,%d-%d", &a, &b, &x, &y)

		if err != nil {
			break
		}

		//  a     b                 a b
		//    x y                 x     y
		if (a <= x && y <= b) || (x <= a && b <= y) {
			p1++
		}

		//    a   b        a   b
		//  x   y            x   y
		if x <= b && a <= y {
			p2++
		}
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
