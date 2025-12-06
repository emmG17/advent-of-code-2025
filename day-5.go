package main

import (
	"slices"
	"fmt"
	"strings"
)

func parseRange(rangeStr string) [][]int {
	var start, end int
	var ranges [][]int
	for r := range strings.SplitSeq(rangeStr, "\n") {
		fmt.Sscanf(r, "%d-%d", &start, &end)
		ranges = append(ranges, []int{start, end})
	}
	return ranges
}

func parseStock(stockStr string) []int {
	var stock []int
	for s := range strings.SplitSeq(stockStr, "\n") {
		var val int
		fmt.Sscanf(s, "%d", &val)
		stock = append(stock, val)
	}
	return stock
}

func DayFive() {
	data := getData("./challenge-input/day-5.txt", "\n\n")
	rangeData := data[0]
	stock := data[1]

	mergedRanges := mergeRanges(parseRange(rangeData))

	part1 := 0
	part2 := 0
	for _, val := range parseStock(stock) {
		if slices.ContainsFunc(mergedRanges, func(r []int) bool {
			return val >= r[0] && val <= r[1]
		}) {
			part1++
		}
	}


	// Count now the number of fresh items (ints in ranges)
	for _, r := range mergedRanges {
		part2 += r[1] - r[0] + 1
	}

	fmt.Printf("Day 5 - Part 1: %d\n", part1)
	fmt.Printf("Day 5 - Part 2: %d\n", part2)
}

func mergeRanges(ranges [][]int) [][]int {
	if len(ranges) == 0 {
		return ranges
	}

	// Sort ranges by their start values
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})

	merged := [][]int{ranges[0]}

	for _, current := range ranges[1:] {
		lastMerged := merged[len(merged)-1]

		if current[0] <= lastMerged[1] {
			// Overlapping ranges, merge them
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// Non-overlapping range, add to merged list
			merged = append(merged, current)
		}
	}

	return merged
}
