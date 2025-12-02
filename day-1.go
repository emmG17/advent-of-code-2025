package main

import (
	"fmt"
	"math"
	"strconv"
)

// Helper to calculate absolute value for ints
func abs(a int) int {
	return int(math.Abs(float64(a)))
}

// Helper for floor division to determine which "cycle" of 100 the number is in
func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func DayOne() {
	// Dial starts at 50
	dial := 50
	// 'a' stores the answer for Part 1 (lands on 0)
	a := 0
	// 'b' stores the answer for Part 2 (passes 0)
	b := 0
	
	lines := getData("./challenge-input/day-1.txt", "\n")

	for _, line := range lines {
		direction := line[0]
		
		// Parse the movement amount
		number := line[1:]
		num, err := strconv.Atoi(number)
		check(err)

		// L moves down/left (negative), R moves up/right (positive)
		if direction == 'L' {
			num = -num
		}

		// Calculate the new position on the infinite number line
		sum := dial + num

		// PART 1: Did we land exactly on a multiple of 100?
		if sum % 100 == 0 {
			a += 1
		}

		// PART 2: Calculate how many 100s we crossed.
		// By subtracting the floor(x/100) of start and end, we find the 
		// number of boundaries crossed.
		rounds := abs(floorDiv(sum, 100) - floorDiv(dial, 100))
		b += rounds

		// PART 2 CORRECTION:
		// Logic adjustment for negative movement (Left turns)
		if (sum < dial) {
			dialC := dial % 100 == 0
			sumC := sum % 100 == 0
			dialI := 0
			sumI := 0

			if dialC { dialI = 1 }
			if sumC { sumI = 1 }
			
			// If we landed on 0 (sumC), the floor logic misses it (add 1).
			// If we started on 0 (dialC), the floor logic counts it (sub 1).
			x := sumI - dialI 
			b += x
		}

		// Update current position for next iteration
		dial = sum
	}

	fmt.Println("Part 1:", a, "Part 2:", b)
}

