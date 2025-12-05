# Advent of Code - Day 4: The Paper Roll Grid

## 1. The Problem
We are in a warehouse full of paper rolls arranged in a grid.
* **Symbol `@`**: A roll of paper.
* **Symbol `.`**: Empty space.

**The Rule:** A forklift can access (and remove) a paper roll if it has **fewer than 4 neighbors**.
(Neighbors include the 8 surrounding cells: horizontal, vertical, and diagonal).

* **Part 1:** How many rolls are currently accessible?
* **Part 2:** If we remove the accessible rolls, *new* rolls might become accessible (because they lost a neighbor). Keep removing them until none are left to remove. How many did we remove in total?

This is a variation of a classic CS concept called **Cellular Automata** (like Conway's Game of Life).

---

## 2. Data Structure: Flattening the Grid
In most intro classes, you learn to represent a grid using a 2D array (e.g., `grid[row][col]`).
However, computers actually store memory in a single long line.

For this solution, we used a **1D Array** to represent the **2D Grid**.



### The Math
To find a cell at `(row, col)` in a 1D array:
$$index = (row \times width) + col$$

**Why do this?**
1.  **Memory:** It allocates one continuous block of memory, which is cache-friendly.
2.  **Simplicity:** We only manage one loop and one slice, rather than a slice of slices.

---

## 3. The Logic: Finding Neighbors
The hardest part of this challenge is correctly identifying the 8 neighbors of a specific cell without crashing the program (going out of bounds) or wrapping around the map incorrectly.



### The Offsets
If we are at index `i` in a 1D array, our neighbors are at fixed "offsets" relative to us.
If the grid is `10` cells wide:
* The neighbor directly **above** is `i - 10`.
* The neighbor directly **below** is `i + 10`.
* The neighbor to the **left** is `i - 1`.

### The Edge Case (Wrap-around Protection)
If we are at the end of a row, `i + 1` technically points to the *start* of the next row in memory. But physically, that cell is not our neighbor—it's on the other side of the room!

We solve this with a coordinate check:
```go
// Calculate row/col for both the current cell and the potential neighbor
row := idx / width
neighborRow := neighborIdx / width

// Only count it if the coordinate distance is <= 1
if abs(row - neighborRow) <= 1 && abs(col - neighborCol) <= 1 {
    // Valid neighbor
}
```

-----

## 4\. The Solution Code

### Part 1: Counting

We scan the grid once. For every `@`, we call `countNeighbors`. If the result is $< 4$, we count it.

### Part 2: The "Naive" Simulation

For Part 2, we use a simple loop that runs until the job is done.

1.  Scan the whole grid.
2.  Remove any accessible paper.
3.  If we removed something, **repeat step 1**.
4.  If we didn't remove anything, **stop**.

<!-- end list -->

```go
package main

import (
	"fmt"
	"strings"
)

// Main logic for Part 1
func countAccessibleRolls(input string) int {
	// ... (Parsing logic) ...
	
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

// Main logic for Part 2
// This is an iterative approach: "Keep scanning until nothing changes"
func countTotalRemovableNaive(input string) int {
    // ... (Parsing logic) ...
	
	totalRemoved := 0
	removedThisIteration := true
	
	// The Loop: Keep going as long as we are making progress
	for removedThisIteration {
		removedThisIteration = false
		
		for i := range grid {
			if grid[i] == '@' {
				neighbors := countNeighbors(grid, i, width)
				if neighbors < 4 {
					// Modify IN PLACE: Remove the paper immediately
					grid[i] = '.' 
					totalRemoved++
					removedThisIteration = true
				}
			}
		}
	}
	return totalRemoved
}

// The Helper: Counts the 8 surrounding neighbors
func countNeighbors(grid []rune, idx, width int) int {
	count := 0
	row := idx / width
	col := idx % width
	
    // The offsets relative to the current position
	offsets := []int{
		-(width + 1), -width, -(width - 1), // Top row
		-1,                    1,           // Left and Right
		(width - 1),  width,  (width + 1),  // Bottom row
	}
	
	for _, offset := range offsets {
		neighborIdx := idx + offset
		
		// 1. Check if index exists (Bounds Check)
		if neighborIdx < 0 || neighborIdx >= len(grid) {
			continue
		}
		
		// 2. Check for "Wrap Around" errors
		neighborRow := neighborIdx / width
		neighborCol := neighborIdx % width
		
		if abs(row-neighborRow) <= 1 && abs(col-neighborCol) <= 1 {
			if grid[neighborIdx] == '@' {
				count++
			}
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 { return -x }
	return x
}
```

-----

## 5\. Engineering Note: Time vs. Space

The solution for Part 2 is labeled "Naive." Why?

**The Problem:**
Imagine a grid with 1,000,000 cells. If we remove **one** piece of paper, we re-scan all 1,000,000 cells again to see if anything changed. This works fine for small inputs, but it is inefficient.

**The Optimization (Trade-off):**
We could use a **Queue**.

1.  When we remove a paper, we know *only its neighbors* are affected.
2.  Add those neighbors to a list (Queue) to check next.
3.  Only check the cells in the Queue.

  * **Naive Approach:** Saves Memory (Space), Costs CPU (Time).
  * **Queue Approach:** Saves CPU (Time), Costs Memory (Space to store the queue).

For this challenge, the grid is small, so the Naive approach is the correct engineering choice: **Simple to write, fast enough to run.**
