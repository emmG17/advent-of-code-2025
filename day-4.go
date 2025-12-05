package main

import (
	"fmt"
	"strings"
)

func countTotalRemovableNaive(input string) int {
	// repeated code from countAccessibleRolls - could be refactored but whatever
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}
	
	height := len(lines)
	width := len(lines[0])
	
	grid := make([]rune, height*width)
	for i, line := range lines {
		for j, char := range line {
			grid[i*width+j] = char
		}
	}
	
	totalRemoved := 0
	removedThisIteration := true
	
	// Keep removing until no more can be removed, this is naive but works
	for removedThisIteration {
		removedThisIteration = false
		
		for i := range grid {
			if grid[i] == '@' {
				neighbors := countNeighbors(grid, i, width)
				if neighbors < 4 {
					grid[i] = '.'
					totalRemoved++
					removedThisIteration = true
				}
			}
		}
	}
	return totalRemoved
}

func countAccessibleRolls(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}
	
	height := len(lines)
	width := len(lines[0])
	
	grid := make([]rune, height*width)
	for i, line := range lines {
		for j, char := range line {
			grid[i*width+j] = char
		}
	}
	
	accessible := 0
	for i := range grid {
		if grid[i] == '@' {
			neighbors := countNeighbors(grid, i, width)
			if neighbors < 4 {
				accessible++
			}
		}
	}
	
	return accessible
}

func countNeighbors(grid []rune, idx, width int) int {
	count := 0
	row := idx / width
	col := idx % width
	
	offsets := []int{
		-(width + 1), -width, -(width - 1),
		-1, 1,
		(width - 1), width, (width + 1),
	}
	
	for _, offset := range offsets {
		neighborIdx := idx + offset
		
		// Bounds check
		if neighborIdx < 0 || neighborIdx >= len(grid) {
			continue
		}
		
		// Calculate neighbor's row and column
		neighborRow := neighborIdx / width
		neighborCol := neighborIdx % width
		
		// Check if truly adjacent (prevent wrapping across rows)
		if abs(row-neighborRow) <= 1 && abs(col-neighborCol) <= 1 {
			if grid[neighborIdx] == '@' {
				count++
			}
		}
	}
	
	return count
}

func DayFour() {
 	input := getData("./challenge-input/day-4.txt", "")
	data := strings.Join(input, "")
	part1 := countAccessibleRolls(data)
	part2 := countTotalRemovableNaive(data)
	fmt.Printf("Day 4 - Part 1: %d\n", part1)
	fmt.Printf("Day 4 - Part 2: %d\n", part2)
}
