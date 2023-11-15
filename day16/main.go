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

// DFS to check different paths for the max score
func maxScore(dist [][]int, rates []int, currentIndex int, visited []bool, timeLeft int, score int, max *int) {
	if score > *max {
		*max = score
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
		maxScore(dist, rates, i, newVisited, newTimeLeft, newScore, max)
	}
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
	p1 := 0
	visited := make([]bool, len(rates))
	visited[startIndex] = true
	maxScore(dist, rates, startIndex, visited, 30, 0, &p1)

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
