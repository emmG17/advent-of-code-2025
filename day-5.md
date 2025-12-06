# Advent of Code - Day 5: The Cafeteria Inventory

## 1. The Problem
We are helping kitchen elves sort ingredients based on ID numbers.
* **Input:** A list of "Fresh" ID ranges (e.g., `3-5`, `10-14`) and a list of specific IDs in stock.
* **Part 1:** Which items in stock fall inside *any* valid fresh range?
* **Part 2:** How many *unique* integer IDs are valid in total, across all ranges?

## 2. The Core Challenge: Overlapping Ranges
The tricky part of this problem is that ranges can **overlap**.
* Range A: `10-20`
* Range B: `15-25`

If we treat these separately in Part 2, we might count the numbers `15, 16, 17...20` twice. This leads to the wrong answer.

### The Solution: Interval Merging
To solve this efficiently, we must simplify the data before we count. We merge overlapping intervals into single, continuous blocks.



**The Algorithm:**
1.  **Sort:** Arrange all ranges by their **Start** number (Low to High).
2.  **Iterate:** Take the first range. Compare it to the next one.
3.  **Merge logic:**
    * If `Next_Start <= Current_End`: They overlap (or touch). Extend `Current_End` to `max(Current_End, Next_End)`.
    * Else: The ranges are separate. Save the current one and start a new block.

**Visualizing the Merge:**

```text
Range 1:  [----------]       (10-20)
Range 2:       [-----------] (15-25)
Merged:   [----------------] (10-25)

Range 3:                       [-----] (30-35)
Final:    [----------------]   [-----]
```

## 3. The Code Solution (Go)
This solution utilizes Modern Go (1.23+) features, specifically iterators (strings.SplitSeq) and the slices package for sorting.

### Helper: Parsing
Note the use of fmt.Sscanf to extract numbers from formatted strings like "3-5".

```go
func parseRange(rangeStr string) [][]int {
	var start, end int
	var ranges [][]int
    // strings.SplitSeq creates an iterator over the lines (Memory efficient!)
	for r := range strings.SplitSeq(rangeStr, "\n") {
		fmt.Sscanf(r, "%d-%d", &start, &end)
		ranges = append(ranges, []int{start, end})
	}
	return ranges
}
```

### The Heavy Lifter: `mergeRanges`
This is the standard implementation of the Interval Merging algorithm.

```go
func mergeRanges(ranges [][]int) [][]int {
	if len(ranges) == 0 {
		return ranges
	}

	// STEP 1: Sort by start time (Crucial!)
    // If we don't sort, we can't merge in a single pass.
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})

	merged := [][]int{ranges[0]}

	// STEP 2: The "Sweep"
	for _, current := range ranges[1:] {
		// Get pointer to the last range we added to our merged list
		lastMerged := merged[len(merged)-1]

		// Check for overlap
		if current[0] <= lastMerged[1] {
			// They overlap. Do we need to extend the end?
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
            // Note: We modify lastMerged in place, no need to append
		} else {
			// No overlap. Start a new range.
			merged = append(merged, current)
		}
	}

	return merged
}
```

### Main Logic
Once the ranges are merged, the math becomes trivial.

```go
func DayFive() {
    // ... Parsing code ...

    // Clean the data first!
	mergedRanges := mergeRanges(parseRange(rangeData))

	part1 := 0
	part2 := 0
    
    // Part 1: Check items against merged ranges
	for _, val := range parseStock(stock) {
		if slices.ContainsFunc(mergedRanges, func(r []int) bool {
			return val >= r[0] && val <= r[1]
		}) {
			part1++
		}
	}

	// Part 2: Calculate total capacity
    // Since ranges don't overlap anymore, we can just sum their lengths.
    // Formula: (End - Start) + 1
    // Example: Range 3-5 contains 3, 4, 5. (5-3) + 1 = 3 items.
	for _, r := range mergedRanges {
		part2 += r[1] - r[0] + 1
	}

	fmt.Printf("Day 5 Results: Part 1: %d, Part 2: %d\n", part1, part2)
}
```

## 4. Engineering Lesson: Data Sanitation
This problem is a classic example of Data Sanitation.

- Naive Approach: You could try to build a boolean array or a hash map representing every valid number.

    - Risk: If the range is 1-1,000,000,000, your program runs out of RAM.

- Engineering Approach: Sanitize the rules (the ranges) first.

    - By running mergeRanges, we converted a messy list of rules into a clean, non-overlapping set.

    - This makes Part 2 a simple addition problem (O(N)) instead of a massive memory allocation problem.
