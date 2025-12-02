package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Finds if the strings has a repeating pattern by bruteforcing all possible lengths
func theHammer(delicateThing string) bool {
outer:
	for n := 1; n <= len(delicateThing)/2; n++ {
		// If the length of the string is not divisible by n, skip tha thing
		if len(delicateThing)%n != 0 {
			continue
		}

		pattern := delicateThing[:n] // The cake slice: 123, 34, 56, etc.

		// Check if the pattern repeats throughout the string but not all the way to the end, why would we want that if the pattern repeats the whole string?
		for i := n; i <= len(delicateThing)-n; i += n {
			// Comparing cake slices, mine's better
			if pattern != delicateThing[i:i+n] {
				// Continue but the outer loop
				continue outer
			}
		}
		// If we reach the end a slice matched, we found a pattern (of many) 
		return true
	}
	return false
}

// I feel sad I bruteforced this one but I'm tired
func DayTwo() {
	ranges := getData("./challenge-input/day-2.txt", ",")
	sum := 0
	sumB := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")

		lower := bounds[0]
		upper := bounds[1]

		lowerInt, err := strconv.Atoi(lower)
		check(err)
		upperInt, err := strconv.Atoi(upper)
		check(err)

		for id := lowerInt; id <= upperInt; id++ {
			strId := strconv.Itoa(id)
			// Odd length IDs are automatically invalid cuz cannot be repeated evenly
			if theHammer(strId) {
				sumB += id
			}

			if len(strId)%2 != 0 {
				continue
			}

			mid := len(strId) / 2

			if strId[:mid] == strId[mid:] {
				sum += id
			}
		}
	}

	fmt.Printf("Day 2:\nPart 1: %d, Part 2: %d\n", sum, sumB)

}
