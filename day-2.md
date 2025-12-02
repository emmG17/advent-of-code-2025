# Advent of Code Solution: The Gift Shop Invalid IDs

## 1. Problem Overview
We are tasked with auditing a database of product IDs for a North Pole gift shop. An elf accidentally entered "invalid" IDs based on specific repeating patterns. We are given a list of ID ranges (e.g., `11-22, 95-115`) and must sum up all invalid IDs found within those ranges.

### The Rules
* **Part 1:** An ID is **invalid** if it consists of a sequence repeated exactly **twice**.
    * *Example:* `123123` (Pattern `123` x 2), `55` (Pattern `5` x 2).
    * *Invalid:* `123123123` (Repeats 3 times).
* **Part 2:** An ID is **invalid** if it consists of a sequence repeated **at least twice**.
    * *Example:* `123123` (x2), `123123123` (x3), `11111` (x5).

---

## 2. The Logic

### Part 1: The Simple Split
Since Part 1 strictly requires the pattern to repeat **exactly twice**, the logic is straightforward:
1.  The string length must be **even**. (You cannot split an odd-length string into two equal integer parts).
2.  Split the string exactly in the middle.
3.  Check if `First Half == Second Half`.

### Part 2: "The Hammer" (Brute Force Pattern Matching)
Part 2 is more complex because we don't know the length of the pattern or how many times it repeats. We use a brute-force approach called `theHammer`.

**How `theHammer` works:**
1.  **Iterate Pattern Lengths ($n$):** We try every possible pattern length from 1 up to half the length of the ID string.
2.  **Divisibility Check:** If the total length of the ID string isn't divisible by $n$, that pattern length is impossible. (e.g., A length 10 string cannot be made of a repeating length 3 pattern).
3.  **Slice & Compare:**
    * Take the first $n$ characters as the **Master Pattern**.
    * Walk through the rest of the string in chunks of size $n$.
    * If any chunk does not match the Master Pattern, stop and try the next length.
    * If we reach the end of the string with all chunks matching, the ID is invalid.

---

## 3. The Code Solution (Go)

Below is the implementation. Note the use of `strings` and `strconv` for manipulating the IDs.

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// theHammer checks if a string is composed of a single pattern 
// repeated 2 or more times.
// Input: "121212" -> Returns: true (Pattern "12")
// Input: "121213" -> Returns: false
func theHammer(delicateThing string) bool {
    // We label the loop 'outer' so we can break/continue the main loop 
    // from inside the nested loop.
outer:
	// 1. Try every possible pattern length (n)
	// We only go up to len/2 because a pattern must repeat at least twice.
	for n := 1; n <= len(delicateThing)/2; n++ {
		
		// 2. Optimization: If total length isn't divisible by n, 
        // it can't be a perfect repetition.
		if len(delicateThing)%n != 0 {
			continue
		}

		pattern := delicateThing[:n] // The "Master Pattern"

		// 3. Verify the pattern repeats across the entire string
		for i := n; i <= len(delicateThing)-n; i += n {
			// Compare the current slice against the Master Pattern
			if pattern != delicateThing[i:i+n] {
				// Mismatch found, this pattern length 'n' is invalid.
				// Continue to the next 'n' in the outer loop.
				continue outer
			}
		}
		
		// If we finish the inner loop without mismatch, we found a valid pattern!
		return true
	}
	return false
}

func DayTwo() {
    // Assumption: getData parses the text file into a slice of strings like "11-22"
	ranges := getData("./challenge-input/day-2.txt", ",")
	
	sumPart1 := 0
	sumPart2 := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")

		// Convert range bounds to integers
		lowerInt, _ := strconv.Atoi(bounds[0])
		upperInt, _ := strconv.Atoi(bounds[1])

		// Check every ID in the range
		for id := lowerInt; id <= upperInt; id++ {
			strId := strconv.Itoa(id)

			// --- PART 2 CHECK ---
			// We check Part 2 first (or separately) using the generic Hammer function
			if theHammer(strId) {
				sumPart2 += id
			}

			// --- PART 1 CHECK ---
			// Optimization: Odd length strings cannot be split evenly
			if len(strId)%2 != 0 {
				continue
			}

			mid := len(strId) / 2
            
            // Split in half and compare
			if strId[:mid] == strId[mid:] {
				sumPart1 += id
			}
		}
	}

	fmt.Printf("Day 2 Results:\nPart 1 Sum: %d\nPart 2 Sum: %d\n", sumPart1, sumPart2)
}
```

## 4. Complexity Analysis
You might wonder if "brute forcing" via theHammer is too slow.Input Size: The IDs are product IDs (integers). Even a large integer (e.g., 2 billion) has only ~10 digits.Iterations: For a 10-digit number, theHammer loop runs for $n=1$ to $n=5$.Comparisons: The inner string comparisons are very cheap for short strings.While the time complexity for checking one string is roughly $O(L^2)$ (where $L$ is the number of digits), $L$ is so small in this context that this approach is highly efficient and readable.
