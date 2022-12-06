package main

import (
	"bufio"
	"fmt"
	"os"
)

func isMarker(freq []int) bool {
	for _, n := range freq {
		if n > 1 {
			return false
		}
	}

	return true
}

func markerIndex(input string, size int) int {
	freq := make([]int, 26)

	for _, c := range input[:size] {
		freq[c-'a']++
	}

	for i, c := range input[size:] {
		if isMarker(freq) {
			return i + size
		}

		freq[input[i]-'a']--
		freq[c-'a']++
	}

	if isMarker(freq) {
		return len(input)
	} else {
		return -1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Println("Part1:", markerIndex(input, 4))
	fmt.Println("Part2:", markerIndex(input, 14))
}
