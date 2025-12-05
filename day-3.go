package main

import (
	"fmt"
	"strconv"
)

func bankJoltage(bank []int, batteries int) int {
		rowLength := len(bank)

		validRangeEnd := (rowLength - batteries) + 1
		currentDigit, relativeIdx := findMaxDigit(bank[:validRangeEnd])

		resultDigits := []int{currentDigit}
		currentIdx := relativeIdx

		// Better go backwards to avoid complicated index management
		for remaining := batteries - 1; remaining > 0; remaining-- {
			currentIdx++

			availableSlice := bank[currentIdx:]
			availableLen := len(availableSlice)

			if remaining == availableLen {
				resultDigits = append(resultDigits, availableSlice...)
				break
			}

			windowSize := (availableLen - remaining) + 1

			bestDigit, offset := findMaxDigit(availableSlice[:windowSize])

			currentIdx += offset
			resultDigits = append(resultDigits, bestDigit)
		}

		currentVal := 0
		for _, d := range resultDigits {
			currentVal = currentVal*10 + d // Sum the digits into a single number, faster than using strings (?!) 
		}
		return currentVal
	}


func LeJoltage(banks []string, batteries int) int {
	sum := 0
	for _, bank := range banks {

		nums := stringToIntSlice(bank)
		
		currentVal := bankJoltage(nums, batteries)

		sum += currentVal
	}
	return sum
}

func DayThree() {
	banks := getData("./challenge-input/day-3.txt", "\n")
	part1 := LeJoltage(banks, 2)
	part2 := LeJoltage(banks, 12)

	fmt.Println("Day 3 complete!")
	fmt.Printf("Part 1 Solution: %d\n", part1)
	fmt.Printf("Part 2 Solution: %d\n", part2)

}

func stringToIntSlice(s string) []int {
	nums := make([]int, len(s))
	for i, char := range s {
		num, _ := strconv.Atoi(string(char))
		nums[i] = num
	}
	return nums
}

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
