package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Files map[string]*File

// Directories are files with children
type File struct {
	Name     string
	Size     int
	Children Files
}

func computeSizes(dir *File) int {
	// It's a file
	if len(dir.Children) == 0 {
		return dir.Size
	}

	size := 0

	for k, file := range dir.Children {
		if k == ".." {
			continue
		}

		size += computeSizes(file)
	}

	dir.Size = size

	return size
}

func dfs(dir *File, total *int, spaceNeeded int, minSize *int) {
	// It's a directory
	if len(dir.Children) > 0 {
		// Part 1
		if dir.Size <= 100000 {
			*total += dir.Size
		}

		// Part 2
		if dir.Size >= spaceNeeded && dir.Size < *minSize {
			*minSize = dir.Size
		}
	}

	for k, file := range dir.Children {
		if k == ".." {
			continue
		}

		dfs(file, total, spaceNeeded, minSize)
	}
}

func solve(lines []string) (int, int) {
	current := &File{Name: "/", Children: Files{}}
	root := current

	// Skip first "$ cd /", we're already in root folder
	for _, line := range lines[1:] {
		switch {
		case line[:4] == "$ cd":
			name := line[5:]
			current = current.Children[name]
		case line[:3] == "dir":
			name := line[4:]
			current.Children[name] = &File{
				Name:     name,
				Children: Files{"..": current},
			}
		case '0' <= line[0] && line[0] <= '9':
			split := strings.Split(line, " ")
			name := split[1]
			size, _ := strconv.Atoi(split[0])
			current.Children[name] = &File{Name: name, Size: size}
		}
	}

	// Add sizes to all dirs in the tree
	computeSizes(root)

	// Part 1
	total := 0

	// Part 2
	freeSpace := 70000000 - root.Size
	spaceNeeded := 30000000 - freeSpace
	minSize := 1<<32 - 1

	// Run one DFS to compute results for both parts
	dfs(root, &total, spaceNeeded, &minSize)

	return total, minSize
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
