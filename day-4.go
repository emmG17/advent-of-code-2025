package main

import (
	"fmt"
	"strings"
)

func countAccessibleRolls(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}
	
	height := len(lines)
	width := len(lines[0])
	
	// Convert 2D grid to 1D array
	grid := make([]rune, height*width)
	for i, line := range lines {
		for j, char := range line {
			grid[i*width+j] = char
		}
	}
	
	// Count accessible rolls
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
	result := countAccessibleRolls(data)
	fmt.Println("Number of accessible rolls:", result)
}
