package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// All pair shortest distances using Floyd–Warshall algorithm
func computeShortestDistances(graph [][]int) [][]int {
	n := len(graph)
	dist := make([][]int, n)
	for i := range graph {
		dist[i] = make([]int, n)
	}

	for i, row := range graph {
		for j, weight := range row {
			if weight > 0 {
				dist[i][j] = weight
			} else {
				dist[i][j] = 100000
			}
		}
	}

	for i := range graph {
		dist[i][i] = 0
	}

	for k := range graph {
		for i := range graph {
			for j := range graph {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

// No point in opening zero rate valves so we can remove them from the
// fully connected graph.
func pruneZeroRates(dist [][]int, rates []int) ([][]int, []int) {
	n := 0
	for _, rate := range rates {
		if rate != 0 {
			n++
		}
	}
	newRates := make([]int, n)
	newDist := make([][]int, n)
	for i := range newDist {
		newDist[i] = make([]int, n)
	}

	x := 0
	y := 0
	for i, row := range dist {
		if rates[i] == 0 {
			continue
		}

		newRates[x] = rates[i]
		y = 0

		for j := range row {
			if rates[j] == 0 {
				continue
			}

			newDist[x][y] = dist[i][j]
			y++
		}

		x++
	}

	return newDist, newRates
}

func toBitstring(visited []bool) int {
	res := 0

	for i, vis := range visited {
		if vis {
			res |= 1 << i
		}
	}

	return res
}

// DFS to check different subpaths and saves them in memo
// memo is 2D array of [endIndex][visitedBitstring] = maxScore
func maxScores(dist [][]int, rates []int, currentIndex int, visited []bool, timeLeft int, score int, memo [][]int) {
	if score > memo[currentIndex][toBitstring(visited)] {
		memo[currentIndex][toBitstring(visited)] = score
	}

	for i, vis := range visited {
		if vis {
			continue
		}

		newTimeLeft := timeLeft - dist[currentIndex][i] - 1
		if newTimeLeft <= 0 {
			continue
		}

		newVisited := make([]bool, len(visited))
		copy(newVisited, visited)
		newVisited[i] = true

		newScore := score + (newTimeLeft * rates[i])
		maxScores(dist, rates, i, newVisited, newTimeLeft, newScore, memo)
	}
}

func computeMaxScores(dist [][]int, rates []int, startIndex int, timeLeft int) [][]int {
	n := len(rates)
	memo := make([][]int, n)
	for i := range rates { // each possible end node has
		memo[i] = make([]int, 1<<n) // 2^n possible subpaths
	}
	visited := make([]bool, n)
	visited[startIndex] = true
	maxScores(dist, rates, startIndex, visited, timeLeft, 0, memo)

	return memo
}

// Search all subpath solutions for the max score
func part1(memo [][]int) int {
	max := 0
	for _, row := range memo {
		for _, score := range row {
			if score > max {
				max = score
			}
		}
	}
	return max
}

// part 2 solution will have 2 subpaths that overlap only
// on the start node
func areDisjoint(visited1 int, visited2 int, startIndex int) bool {
	return (visited1 & visited2) == (1 << startIndex)
}

// Search pairs of subpath solutions that have disjoint visited nodes
// and find the max pair score
func part2(memo [][]int, startIndex int) int {
	n := len(memo)
	max := 0

	for i := 0; i < n; i++ {
		for visited1, score1 := range memo[i] {
			// the solution must have 2 non-zero scores
			if score1 == 0 {
				continue
			}

			for j := i + 1; j < n; j++ {
				for visited2, score2 := range memo[j] {
					if !areDisjoint(visited1, visited2, startIndex) {
						continue
					}

					if score1+score2 > max {
						max = score1 + score2
					}
				}
			}
		}
	}

	return max
}

func solve(lines []string) (int, int) {
	n := len(lines)
	nodeIndex := make(map[string]int, n)
	rates := make([]int, n)

	for i, line := range lines {
		split := strings.Split(line, ";")
		valvePart := split[0]
		var mainId string
		var rate int
		fmt.Sscanf(valvePart, "Valve %s has flow rate=%d", &mainId, &rate)
		nodeIndex[mainId] = i
		if mainId == "AA" {
			rates[i] = -1
		} else {
			rates[i] = rate
		}
	}

	graph := make([][]int, n)
	for i := range lines {
		graph[i] = make([]int, n)
	}

	for i, line := range lines {
		split := strings.Split(line, ";")
		ids := strings.Split(split[1], ", ")
		ids[0] = ids[0][len(ids[0])-2:] // remove prefix text

		for _, id := range ids {
			j := nodeIndex[id]
			graph[i][j] = 1
		}
	}

	dist := computeShortestDistances(graph)
	dist, rates = pruneZeroRates(dist, rates)
	var startIndex int
	for i, rate := range rates {
		if rate == -1 {
			startIndex = i
			break
		}
	}

	// Part 1
	memo := computeMaxScores(dist, rates, startIndex, 30)
	p1 := part1(memo)

	// Part 2
	memo = computeMaxScores(dist, rates, startIndex, 26)
	p2 := part2(memo, startIndex)

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
