# Advent of Code - Day 3: The Escalator Batteries

## 1. The Problem
We are stuck in a lobby because the elevators and escalators are out of power. To fix the escalator, we need to configure banks of batteries.

* **Input:** We are given several rows of numbers (battery banks).
    * Example: `987654321111111`
* **The Goal:** We need to select a specific number of batteries ($k$) from the bank to create the **largest possible number**.
* **The Constraints:**
    1.  **Order Matters:** We cannot rearrange the batteries. If the input is `12345`, we can pick `1` and `5` to make `15`, but we cannot make `51`.
    2.  **Part 1:** We must pick exactly **2** batteries.
    3.  **Part 2:** We must pick exactly **12** batteries.

---

## 2. The Logic: "The Greedy Window"

The intuitive solution might be to "just pick the biggest number." But we have to be careful. If we pick a number too far to the right, we might run out of room to pick the rest of the required batteries.

### The Strategy
We use a **Greedy Approach**. This means at every step, we try to make the best possible choice *right now*, assuming it will lead to the best result overall.

Since we want the largest total number, we want the **largest digit possible in the first position** (the left-most side). Then we want the largest possible in the second position, and so on.

### The "Safety Window"
Imagine we have the row: `4 9 1 2 5` and we need to pick **3** numbers.


**Step 1:** We need to pick our first number.
* Can we pick the `5` at the end? **No.** If we pick `5`, there are 0 digits left after it, but we still need to pick 2 more numbers.
* We must leave "buffer space" for the remaining digits.
* Since we need 2 more digits after this one, we must stop looking 2 spots from the end.
* **Search Window:** `[4 9 1] 2 5`
* **Action:** The largest number in the window is `9`. We pick it.

**Step 2:** We have `9`. We need **2** more numbers.
* We start searching *after* the `9`. Remaining list: `1 2 5`.
* We need 1 more digit after this current pick. So we must leave 1 spot buffer.
* **Search Window:** `[1 2] 5`
* **Action:** The largest is `2`. We pick it.

**Step 3:** We have `92`. We need **1** more number.
* Start searching after the `2`. Remaining list: `5`.
* **Search Window:** `[5]`
* **Action:** Pick `5`.

**Final Result:** `925`

---

## 3. The Code Solution (Go)

The solution uses two main functions. One handles the "search window" logic (`bankJoltage`), and the other converts strings to integers (`LeJoltage`).

### Key Variables in `bankJoltage`
* `batteries`: How many digits we still need to find.
* `validRangeEnd`: The limit of how far right we can look without running out of buffer space.
* `offset`: How far into the slice the best digit was found (so we can move our starting point forward).

```go
package main

import (
	"fmt"
	"strconv"
)

// bankJoltage calculates the maximum number possible by picking 'batteries' count of digits.
func bankJoltage(bank []int, batteries int) int {
    rowLength := len(bank)

    // --- STEP 1: Find the very first digit ---
    // We calculate the safety margin. If we need 12 batteries, 
    // we can't look at the last 11 items for our first pick.
    validRangeEnd := (rowLength - batteries) + 1
    
    // findMaxDigit returns the value and its index (relative to the slice we gave it)
    currentDigit, relativeIdx := findMaxDigit(bank[:validRangeEnd])

    // Start our result list
    resultDigits := []int{currentDigit}
    
    // currentIdx tracks where we are in the real array (absolute index)
    currentIdx := relativeIdx

    // --- STEP 2: Find the remaining digits ---
    // We loop backwards from (batteries-1) down to 1
    for remaining := batteries - 1; remaining > 0; remaining-- {
        // Move our pointer one step past the last digit we picked
        currentIdx++

        // Look at everything available to the right
        availableSlice := bank[currentIdx:]
        availableLen := len(availableSlice)

        // Optimization: If the number of items left is exactly what we need,
        // just take them all! No need to search.
        if remaining == availableLen {
            resultDigits = append(resultDigits, availableSlice...)
            break
        }

        // Calculate the new window size based on how many we still need
        windowSize := (availableLen - remaining) + 1

        // Find the best digit in this new window
        bestDigit, offset := findMaxDigit(availableSlice[:windowSize])

        // Advance our absolute pointer by the offset where we found the digit
        currentIdx += offset
        resultDigits = append(resultDigits, bestDigit)
    }

    // --- STEP 3: Combine digits into one number ---
    // Example: [9, 8, 3] -> 983
    currentVal := 0
    for _, d := range resultDigits {
        // Mathematical trick to append a digit: 
        // Multiply total by 10 (shift left) and add the new digit.
        currentVal = currentVal*10 + d 
    }
    return currentVal
}

// Helper to find the largest number and its index in a slice
func findMaxDigit(data []int) (int, int) {
    maxVal, maxIdx := -1, 0
    for i, val := range data {
        if val > maxVal {
            maxVal = val
            maxIdx = i
        }
    }
    return maxVal, maxIdx
}

// Main logic wrapper
func LeJoltage(banks []string, batteries int) int {
    sum := 0
    for _, bank := range banks {
        // Convert the string "123" into a slice of ints [1, 2, 3]
        nums := stringToIntSlice(bank)
        
        // Run the logic
        currentVal := bankJoltage(nums, batteries)

        sum += currentVal
    }
    return sum
}

// Standard data loading and execution
func DayThree() {
    banks := getData("./challenge-input/day-3.txt", "\n")
    
    // Part 1: Pick 2 batteries
    part1 := LeJoltage(banks, 2)
    
    // Part 2: Pick 12 batteries
    part2 := LeJoltage(banks, 12)

    fmt.Println("Day 3 complete!")
    fmt.Printf("Part 1 Solution: %d\n", part1)
    fmt.Printf("Part 2 Solution: %d\n", part2)
}

// Helper to parse input
func stringToIntSlice(s string) []int {
    nums := make([]int, len(s))
    for i, char := range s {
        num, _ := strconv.Atoi(string(char))
        nums[i] = num
    }
    return nums
}
```

-----

## 4\. Engineering Lesson: "You Ain't Gonna Need It" (YAGNI)

In Computer Science, there is often a temptation to build the "perfect" machine—a solution that handles every possible case, creates generic abstractions, and prepares for future features that haven't been asked for yet.

**Don't do this.**

Day 3 is a perfect example of why you should solve the problem *in front of you*, not the problem you *imagine* might happen later.

### Iteration 1: Solving for Part 1

When we started, the requirement was simple: **Find exactly 2 batteries.**

Our first solution was hard-coded for this specific requirement. Look at how simple the logic is:

```go
func DayThreePart1() {
	banks := getData("./challenge-input/day-3.txt", "\n")

	sum := 0
	for j, bank := range banks {
		maxA := 0 // The first digit
		maxB := 0 // The second digit
		
        // Simple linear scan
		for i := 0; i < len(bank) - 1; i++ {
			numA, _ := strconv.Atoi(string(bank[i]))
			numB, _ := strconv.Atoi(string(bank[i + 1]))

            // Hardcoded logic for "The First Battery"
			if maxA < numA {
				maxA = numA
				maxB = 0 // Reset B because we moved A forward
			}

            // Hardcoded logic for "The Second Battery"
			if maxB < numB {
				maxB = numB
			}
		}
		sum += maxA * 10 + maxB
	}
    // ...
}
```

**Why this is good code (for Part 1):**

1.  **Readability:** You can read it top-to-bottom. There is no recursion, no dynamic window calculation, and no slice manipulation.
2.  **Variables are explicit:** We literally named them `maxA` and `maxB`. It is obvious what they do.
3.  **Speed to Ship:** This code takes 5 minutes to write.

### Iteration 2: The Generalized Solution

Then, **Part 2** arrived. The requirement changed from "Find 2" to "Find 12."
Suddenly, `maxA` and `maxB` aren't enough. We would need `maxC`, `maxD`, `maxE`... all the way to `maxL`. The hardcoded `if` statements would become a nightmare.

*Now*—and only now—is the time to engineer the complex **General Solution** we documented above (`bankJoltage`).

### The Trade-off

Compare the complexity:

| Feature | Iteration 1 (Simple) | Iteration 2 (Generalized) |
| :--- | :--- | :--- |
| **Variables** | `maxA`, `maxB` (Integers) | `resultDigits` (Slice), `windowSize`, `offset` |
| **Logic** | Single `for` loop | Nested loops with dynamic resizing |
| **Cognitive Load** | Low (Easy to understand) | High (Requires visualizing the "Window") |
| **Flexibility** | None (Breaks if asked for 3 items) | Infinite (Handles any N items) |

### The Takeaway

If you had started Day 3 trying to write the "Generalized Greedy Window Algorithm" before you even knew Part 2 existed, you would have spent 30 minutes debugging complex slice indices for a problem that could have been solved with two `int` variables.

**Start simple.** Refactor into complexity only when the requirements force you to.
